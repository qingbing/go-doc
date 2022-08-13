package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go test21()
	go test22()
	wg.Wait()

}

func callNum() {
	count := 0
	for i := 0; i < 10000000000; i++ {
		count += i
	}
}
func test21() {
	runtime.LockOSThread()
	start := time.Now()
	callNum()
	end := time.Now()
	fmt.Println("test21 耗时", end.Sub(start))
	fmt.Println()
	defer runtime.UnlockOSThread()
	wg.Done()
}
func test22() {
	start := time.Now()
	callNum()
	end := time.Now()
	fmt.Println("test22 耗时", end.Sub(start))
	fmt.Println(end.Sub(start))
	wg.Done()
}
