package main

import (
	"fmt"
	"math"
)

func main() {
	//输出
	fmt.Println("hello world")
	//整数
	var a = 3
	//浮点数
	var b = 3.14
	var x float32
	var y float64
	x = 3.213
	y = 2.1231

	//字符串
	var c = "hello c"

	//常量
	const h = 500000000
	const i = 3e20 / h
	const st string = "go study time"

	fmt.Println(a, b, c, x, y, math.Sin(i), st)

	// 简写，声明并且赋值
	aa := 123
	bb := "hello bb"
	cc := 3.34123

	dd := "welcome" + bb
	fmt.Println(aa, bb, cc, dd)

	//循环
	for i := 7; i < 9; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//if语句
	if 7%2 == 1 {
		fmt.Println("right")
	} else {
		fmt.Println("false")
	}

	if num := 9; num < 7 {
		fmt.Println("num < 7")
	} else if num < 10 {
		fmt.Println("num < 10")
	} else {
		fmt.Println("num >= 10")
	}

	//switch语句
	switchNum := 2
	switch switchNum {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	//array,长度不可变
	var arr [5]int
	arr[4] = 100
	fmt.Println("get:", arr[4])
	fmt.Println("len:", len(arr))
	args := [5]int{1, 2, 3, 4, 5}
	fmt.Println(args)

	var towA [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(towA[i][j])
		}
	}

	//slice 可以自由变换长度
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)
	copyS := make([]string, len(s))
	copy(copyS, s)
	fmt.Println(copyS)
	fmt.Println(s[:3])
	fmt.Println(s[2:3])

	//map
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	m2 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)
	delete(m2, "one")
	fmt.Println(m2)

	//range
	nums := []int{1, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println(sum)

	mm := map[string]string{"a": "A", "b": "B"}
	for k, v := range mm {
		fmt.Println(k, v)
	}

}
