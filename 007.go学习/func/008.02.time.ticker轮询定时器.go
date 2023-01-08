package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	// 创建定时器
	timer := time.NewTicker(time.Second)

	time.AfterFunc(time.Second*3, func() {
		fmt.Println("Reset")
		timer.Reset(time.Second * 2)
	})
	time.AfterFunc(time.Second*8, func() {
		fmt.Println("Stop")
		timer.Stop()
	})

	go func() {
		for {
			select {
			case now := <-timer.C:
				fmt.Println("Sub Goroutine: ", now)
			case <-stop:
				fmt.Println("Sub Goroutine over:", time.Now())
				return
			}
		}
	}()

	<-time.After(time.Second * 9)
}
