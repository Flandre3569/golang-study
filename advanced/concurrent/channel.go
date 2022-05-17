package main

func CalSquare() {
	src := make(chan int)     // 无缓冲通道 -> 同步通道
	dest := make(chan int, 3) // 有缓冲通道 -> 生产消费模型，如果格子满后会产生阻塞

	// 生产者
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	// 消费者
	// 使用带缓冲的channel，保证生产和消费的速度均衡
	// 防止消费快于生产
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		println(i)
	}
}

func main() {
	CalSquare()
}
