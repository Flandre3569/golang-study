package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

//查询用户，如果查到将用户返回回去，如果没查到，返回错误
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {

	u, err := findUser([]user{{name: "张三", password: "123"}}, "张三")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)

	if u, err := findUser([]user{{"wang", "123"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u)
	}
}
