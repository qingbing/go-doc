package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start: ", time.Now())
	// 第一种定时器: time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("第一种定时器, time.Sleep: ", time.Now())
	// 第二种定时器
	chT := time.After(time.Second)
	fmt.Println("第二种定时器, time.After:", <-chT)

	// 第三种定时器: time.NewTimer(time.Second)
	timer := time.NewTimer(time.Second * 5)
	timer.Reset(time.Second * 2)
	// timer.Stop() // 终止定时器
	fmt.Println("第三种定时器, time.NewTimer:", <-timer.C)
}
