package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	// 创建定时器
	timer := time.NewTicker(time.Second)

	time.AfterFunc(time.Second*4, func() {
		timer.Stop()
		fmt.Println("Stop")
		stop <- struct{}{}
	})

	go func() {
		// 循环读取一旦信号处于 stop 后会一致处于阻塞状态
		for now := range timer.C {
			fmt.Println("当前时间: ", now)
		}
		fmt.Println("hahae")
	}()

	<-stop
	fmt.Println("Program is over")
}
