package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("系统: ", runtime.GOOS)
	fmt.Println("Go Version: ", runtime.Version())
	fmt.Println("Go Path: ", runtime.GOROOT())
	fmt.Println("当前电脑cpu核数: ", runtime.NumCPU())
	n := runtime.GOMAXPROCS(1) // 设置为单核时，需要cpu时间抢占，一个协程执行cpu时间滴答完成后才会进行轮换
	// n := runtime.GOMAXPROCS(2) // 设置会双核，会两个cpu同时执行
	fmt.Println("进程默认使用的CPU核数: ", n)
	for i := 0; i < 1000; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
