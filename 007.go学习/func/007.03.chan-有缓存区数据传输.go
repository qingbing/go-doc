package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // 存储未满3个元素之前，不会阻塞
	fmt.Println("len(ch) = ", len(ch), "cap(ch)=", cap(ch))

	go func() {
		for i := 0; i < 7; i++ {
			ch <- i
			fmt.Println("子go程, i=", i, ", len(ch) = ", len(ch), "cap(ch)=", cap(ch))
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 7; i++ {
		num := <-ch
		fmt.Println("主go程读取: ", num)
	}
	/*
		for num := range ch {
			fmt.Println("主go程读取: ", num)
		}
	*/
}
