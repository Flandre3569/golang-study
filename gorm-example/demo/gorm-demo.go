package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	name     string
	password string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db2019210745?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	rows, err := db.Query("SELECT `name`,`password` FROM users WHERE `id` = ?", 2)
	if err != nil {
		fmt.Println("查询失败")
	}

	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.name, &user.password)

		if err != nil {
			fmt.Println("查找失败")
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		fmt.Println("查找失败")
	}

	fmt.Println(users)

}
