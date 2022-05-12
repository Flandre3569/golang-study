package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 打印当前时间
	t := time.Date(2022, 1, 1, 1, 30, 30, 0, time.UTC)
	t2 := time.Date(2022, 1, 2, 1, 30, 30, 0, time.UTC)
	fmt.Println(t) // 自定义时间
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(t.Format("2006-01-01 15:01:05"))

	diff := t2.Sub(t)
	fmt.Println(diff)
	fmt.Println(diff.Hours(), diff.Minutes(), diff.Seconds())

	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)
	fmt.Println(now.Unix())
}
