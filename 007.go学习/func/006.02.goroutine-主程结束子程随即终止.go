package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("sing")
		time.Sleep(time.Millisecond * 50)
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("dance")
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	go sing()
	go dance()
	fmt.Println("主程终止，之上的goroutine也就随即终止(看不到运行结果)")
}
