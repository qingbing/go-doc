package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 3; i < 8; i++ {
			ch <- i
		}
		close(ch) // 关闭 channel
		// ch <- 8 // send on closed channel
		fmt.Println("子go程结束")
	}()
	for {
		if num, ok := <-ch; ok {
			fmt.Println("读到数据:", num)
		} else {
			fmt.Println("读取完毕")
			fmt.Println("channel 关闭后读取结果为零值:", <-ch)
			fmt.Println("channel 关闭后读取结果为零值:", <-ch)
			fmt.Println("channel 关闭后读取结果为零值:", <-ch)
			break
		}
	}
}
