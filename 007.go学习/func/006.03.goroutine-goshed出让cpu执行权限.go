package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("sing")
		}
	}()
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 出让当前 cpu 时间片
		fmt.Println("this is main dance")
	}
}
