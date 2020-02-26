package models

import (
    "log"
    "fmt"

    "github.com/jinzhu/gorm"
    //只执行init()函数。这个时候就可以使用 import _
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "rwplus-backend/pkg/setting"
)

var db *gorm.DB

// 定义一个通用的结构体，将通用的id，潜入到其他的结构体里
type Model struct {
    ID int `gorm: "primary_key" json:"id"`
}

func init() {
    var (
        err error
        dbType, dbName, user, passWord, host, tablePrefix string
    )

    // 获取app.ini配置文件下database分区所有的key/value
    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatalf("无法获取'database'分区: %v", err)
    }

    dbType = sec.Key("TYPE").String()
    dbName = sec.Key("NAME").String()
    user = sec.Key("USER").String()
    passWord = sec.Key("PASSWORD").String()
    host = sec.Key("HOST").String()
    tablePrefix = sec.Key("TABLE_PREFIX").String()

    db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user,
        passWord,
        host,
        dbName))

    if err != nil {
        log.Println(err)
    }

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
        return tablePrefix + defaultTableName
    }

    // SingularTable单数表名，plural复数表名
    db.SingularTable(true)
    // 设置空闲连接数
    db.DB().SetMaxIdleConns(10)
    // 设置最大连接数
    db.DB().SetMaxOpenConns(100)
}


func CloseDB() {
    defer db.Close()
}