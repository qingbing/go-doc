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
	for {
		time.Sleep(time.Second)
	}
}
