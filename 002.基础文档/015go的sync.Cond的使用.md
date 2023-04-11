# go 语言 sync.WaitGroup 的使用

- [go 语言 sync.WaitGroup 的使用](#go-语言-syncwaitgroup-的使用)
  - [1. 说明](#1-说明)
    - [1.1 sync.Cond 作用](#11-synccond-作用)
    - [1.2 sync.Cond 应用场景](#12-synccond-应用场景)
    - [1.3 sync.Cond 注意事项](#13-synccond-注意事项)
  - [2. 演示示例](#2-演示示例)
    - [2.1 sync.Cond 实现一发多收(广播)](#21-synccond-实现一发多收广播)
    - [2.2 sync.Cond 实现队列](#22-synccond-实现队列)

## 1. 说明

### 1.1 sync.Cond 作用

1. 设置一组协程根据条件阻塞，可以根据不同的条件阻塞
2. 可以根据条件唤醒相对应的协程


### 1.2 sync.Cond 应用场景

1. 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备的情况

### 1.3 sync.Cond 注意事项

1. 被叫方必须持有锁
2. 主叫方可以持有锁，但允许不持有
3. 尽可能的减少无效唤醒

## 2. 演示示例

### 2.1 sync.Cond 实现一发多收(广播)

- demo

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func initList(list *[]int, c *sync.Cond) {
	// 主叫方, 可以持有锁,也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()

	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	// 唤醒所有条件等待的协程
	c.Broadcast()

	// 唤醒某一个 c.Signal()
	// c.Signal()
	// c.Signal()
}

func readList(list *[]int, c *sync.Cond) {
	// 被叫方,必须持有锁
	c.L.Lock()
	defer c.L.Unlock()

	for len(*list) == 0 { //
		fmt.Println("readlist wait")
		c.Wait() // 条件阻塞
	}
	fmt.Println("list data:", *list)
}

// sync.Cond 实现一发多收(广播)
func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})

	go readList(&list, cond)
	go readList(&list, cond)
	go readList(&list, cond)
	time.Sleep(time.Second)
	initList(&list, cond)
}

func main() {
	CondCase()

	time.Sleep(time.Second * 3)
}
```

- output

```text
readlist wait
readlist wait
readlist wait
list data: [0 1 2 3 4 5 6 7 8 9]
list data: [0 1 2 3 4 5 6 7 8 9]
list data: [0 1 2 3 4 5 6 7 8 9]
```

### 2.2 sync.Cond 实现队列

- demo

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type queue struct {
	list []int
	cond *sync.Cond
}

func newQueue() *queue {
	// 初始化队列
	return &queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (q *queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.list = append(q.list, item)
	// 当数据写入成功后,唤醒一个协程来处理数据
	q.cond.Signal()
}

func (q *queue) GetMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.list) < n { // 队列的数量小于 n 时,需要等待
		q.cond.Wait()
	}

	list := q.list[:n]
	q.list = q.list[n:]
	return list
}

// sync.Cond 实现队列
func CondQueueCase() {
	q := newQueue()
	var wg sync.WaitGroup
	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// list := q.GetMany(n)
			list := q.GetMany(10)
			fmt.Printf("%d: %v\n", n, list)
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
}

func main() {
	CondQueueCase()

	time.Sleep(time.Second * 3)
}
```

- output

```text
3: [0 1 2 3 4 5 6 7 8 9]
9: [20 21 22 23 24 25 26 27 28 29]
8: [10 11 12 13 14 15 16 17 18 19]
1: [30 31 32 33 34 35 36 37 38 39]
0: [50 51 52 53 54 55 56 57 58 59]
2: [60 61 62 63 64 65 66 67 68 69]
4: [40 41 42 43 44 45 46 47 48 49]
6: [70 71 72 73 74 75 76 77 78 79]
7: [80 81 82 83 84 85 86 87 88 89]
5: [90 91 92 93 94 95 96 97 98 99]
```