package main

import "fmt"

func hello(i int) {
	println("hello goroutine:" + fmt.Sprint(i))
}

// 并发打印数字0-4，速度更快
func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
}
func main() {
	HelloGoRoutine()
}
