package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	// 初始化一个 os.Signal 类型的channel, 必须使用缓冲通道，否则在信号发送时如果没有准备好接收信号，就会有丢失信号的风险
	signalChan := make(chan os.Signal, 1)
	fmt.Println("program start:")
	// 使用 os.Notify 监听信号
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for ch := range signalChan {
			switch ch {
			case os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT:
				exitFunc(ch, "信号结束")
			}
		}
	}()

	signal.Ignore(syscall.SIGTRAP)
	//signal.Ignore(os.Interrupt) // 一旦设置，ctrl+c 将不能终止程序
	fmt.Printf("os.Kill: %t\n", signal.Ignored(os.Kill))
	fmt.Printf("syscall.SIGTRAP: %t\n", signal.Ignored(syscall.SIGTRAP))

	//signal.Reset()

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("停止监听")
		signal.Stop(signalChan)
	})

	time.AfterFunc(time.Second*10, func() {
		exitFunc(nil, "自然结束")
		stopCh <- struct{}{}
	})
	<-stopCh
}

func exitFunc(ch os.Signal, msg string) {
	fmt.Println("program end:", msg)
	if ch != nil {
		fmt.Println("信号:", ch)
	}
	os.Exit(0)
}
