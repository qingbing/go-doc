package main

import "fmt"

/*
chan-单向channel 的使用
生产者-消费者模式
*/
// 消费者 读 channel
func consumer(ch <-chan int) {
	for iVal := range ch {
		fmt.Println("chan读到数据", iVal)
	}
}

// 生产者 写channel
func producer(ch chan<- int) {
	for i := 2; i < 10; i++ {
		ch <- i
		fmt.Println("写入chan", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 5) // 双向 channel

	go producer(ch) // 写: 双向转写

	consumer(ch) // 读: 双向转读
}
