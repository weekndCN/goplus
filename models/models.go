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

func Setup() {
    var err error

    db, err = gorm.Open(setting.DatabaseSetting.Type,
        fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

    if err != nil {
        log.Fatalf("models.Setup err: %v", err)
    }

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
        return setting.DatabaseSetting.TablePrefix + defaultTableName
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