package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dataCh := make(chan string)
	stop := make(chan struct{})

	go func() {
		// 监听通道 dataCh， 3秒没有数据就退出
		select {
		case data := <-dataCh:
			fmt.Println("Data: ", data)
		case <-time.After(time.Second * 3):
			fmt.Println("data waiting is timeout")
		}
		// 通知结束
		stop <- struct{}{}
	}()

	go func() {
		fmt.Println(time.Now())
		// 获取随机随眠时间
		rand.Seed(int64(time.Now().Nanosecond()))
		sec := rand.Intn(6)
		fmt.Println("Sleep Second: ", sec)
		time.Sleep(time.Second * time.Duration(sec))
		// 向通道发送数据
		dataCh <- fmt.Sprint(time.Now())
	}()

	<-stop
}
