package main

import (
	"fmt"
	"time"
)

type Printer struct {
	ch chan string
}

func (cp *Printer) push(str string) {
	cp.ch <- str
}

func (cp *Printer) run() {
	go func() {
		for printStr := range cp.ch {
			fmt.Println(printStr)
		}
	}()
}

var printer Printer = Printer{
	ch: make(chan string, 5),
}

func init() {
	printer.run()
}

func main() {
	printer.push("Start")
	printer.push("Over")
	time.Sleep(time.Second)
}
