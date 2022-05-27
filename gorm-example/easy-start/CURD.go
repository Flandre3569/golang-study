package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 注意此处实体的key必须是大写开头

type User struct {
	Id       int
	Name     string
	Password string
}

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db2019210745?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		fmt.Println("连接数据库失败")
	}

	// 插入一条内容到数据库
	//user := User{Name: "小红", Password: "123321"}
	//result := db.Create(&user)
	// 返回影响数据库的行数
	//fmt.Println(result.RowsAffected)

	// 查询用户
	var user User
	//db.First(&user, 14)
	//fmt.Println(user)

	// 删除一条信息
	//db.First(&user, 21).Delete(&user)

	// 修改一条信息
	db.First(&user, 14).Update("password", "54250")

}
