package models

import (
    //"time"
    //"fmt"
    //"github.com/jinzhu/gorm"
)

type User struct {
    // 嵌套入models.go定义的Model结构体,用户主键ID
    Model
    Name string `json: "name"`
    Password string `json: "password"`
    Phone string `json:"phone"`
    CreatedDate int `json:"created_date"`
}


// 根据用户名检索用户是否存在
func ExistUserByName(name string) bool {
    var user User
    // 等同select id from user where name=? limit 1;
    db.Select("id").Where("name = ?", name).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}

// 根据手机号检索用户是否存在
func ExistUserByPhone(phone int) bool {
    var user User
    db.Select("id").Where("phone = ?", phone).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}
