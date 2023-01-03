package main

import (
	"fmt"
	"strconv"
)

/*
chan-模拟生产订单的生产者-消费者
*/
type orderInfo struct {
	id   int
	desc string
}

// 订单信息格式化成字符串
func orderDesc(info orderInfo) string {
	return fmt.Sprintf("id:%d, desc: %s", info.id, info.desc)
}

// 生产订单
func producer(ch chan<- orderInfo) {
	for i := 10; i < 22; i++ {
		orderInfo := orderInfo{
			id:   i,
			desc: "Good-" + strconv.Itoa(i),
		}
		ch <- orderInfo
		fmt.Println("新增订单:", orderDesc(orderInfo))
	}
	close(ch)
}

// 消费订单
func consumer(ch <-chan orderInfo) {
	for orderInfo := range ch {
		fmt.Println("处理订单:", orderDesc(orderInfo))
	}
}

func main() {
	ch := make(chan orderInfo, 5)
	go producer(ch)
	consumer(ch)
}
