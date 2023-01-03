package main

import (
	"fmt"
	"time"
)

func printer(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
		time.Sleep(time.Millisecond * 200)
	}
}

var ch chan struct{} = make(chan struct{})

func person1() {
	// 控制打印在前
	printer("Hello")
	<-ch
}

func person2() {
	ch <- struct{}{}
	// 控制打印在后
	printer("World")
}

func main() {
	go person2()
	go person1()
	time.Sleep(time.Second * 3)
}
