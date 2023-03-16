package main

import (
	"fmt"
	"time"
)

func main() {
	waitchan := make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("count --> ", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		panic("Panic.....")
	}()
	fmt.Println("the main over")
	<-waitchan
}
