package main

import "fmt"

func main() {
	var a = 2
	var b = 3
	result := calculate(a, b)
	fmt.Println(result)
	return
}

func calculate(x, y int) int {
	z := x * y
	return z
}
