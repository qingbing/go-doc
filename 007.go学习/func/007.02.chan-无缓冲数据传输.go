package main

import "fmt"

// chan-无缓冲的数据传输
func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i = ", i, "len(ch) = ", len(ch), "cap(ch)=", cap(ch))
		}
		// 通过 chan 通知主程
		ch <- "子go程打印完毕"
	}()

	fmt.Println(<-ch)
}
