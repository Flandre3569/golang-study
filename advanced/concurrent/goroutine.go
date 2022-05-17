package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	println("hello goroutine:" + fmt.Sprint(i))
}

// 使用goroutine -> 协程实现高并发操作
// 并发打印数字0-4，速度更快

func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}

// 优化并发执行策略

func ManyGoWait() {
	var wg sync.WaitGroup // 实现并发安全操作和协程间的同步
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

func main() {
	//HelloGoRoutine()
	ManyGoWait()
}
