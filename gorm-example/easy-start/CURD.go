package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	name     string
	password string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db2019210745"))

	if err != nil {
		fmt.Println("连接数据库失败")
	}
	//var users []User
	//err = db.Select("name", "password").Find(&users, 1).Error
	//
	//if err != nil {
	//	fmt.Println("连接数据库失败")
	//}
	//fmt.Println(users)
	user := User{name: "小红", password: "123321"}
	result := db.Create(&user)
	fmt.Println(result)
}
