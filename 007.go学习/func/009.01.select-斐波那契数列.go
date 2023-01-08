package main

import (
	"fmt"
	"runtime"
)

/*
斐波那契数列
  f(n) = f(n-1) + f(n-2)
  f(1)=f(2)=1
select-斐波那契数列
*/
func fibonacci(data <-chan int, quit <-chan struct{}) {
	for {
		select {
		case d := <-data:
			fmt.Print(d, " ")
		case <-quit:
			fmt.Println("\nfucntion over")
			runtime.Goexit()
		}
	}
}

func main() {
	data := make(chan int)
	quit := make(chan struct{})
	go fibonacci(data, quit)
	x, y := 1, 1
	for i := 0; i < 20; i++ {
		data <- x
		x, y = y, x+y
	}
	quit <- struct{}{}
}
