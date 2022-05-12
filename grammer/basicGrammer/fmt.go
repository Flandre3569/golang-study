package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	a := "hello"
	n := 123
	p := point{x: 1, y: 2}

	fmt.Println(a, n)
	fmt.Println(p)

	fmt.Printf("a=%v\n", a)
	fmt.Printf("p=%v\n", p)  // 打印value
	fmt.Printf("p=%+v\n", p) // 打印key value的形式
	fmt.Printf("a=%#v\n", a) // "hello" 注明类型
	fmt.Printf("p=%#v\n", p) // main.point{x:1, y:2} 注明类型

	f := 3.141592
	fmt.Printf("%.2f\n", f) // 格式化保留小数点位数
}
