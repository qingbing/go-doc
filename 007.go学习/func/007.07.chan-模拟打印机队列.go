package main

import (
	"context"
	"fmt"
	"strconv"
)

/*
chan - 模拟打印机功能
*/
type Printer struct {
	ch     chan string
	ctx    context.Context
	cancel context.CancelFunc
}

func (printer *Printer) start() {
	printer.ctx, printer.cancel = context.WithCancel(context.Background())
	go func() {
		for printStr := range printer.ch {
			fmt.Println(printStr)
		}
		printer.cancel()
	}()
}

func (printer *Printer) close() {
	close(printer.ch)
	<-printer.ctx.Done()
}

func (cp *Printer) push(str string) {
	cp.ch <- str
}

func NewPrinter(num int) *Printer {
	p := &Printer{
		ch: make(chan string, num),
	}
	p.start()
	return p
}

func main() {
	printer := NewPrinter(5)
	for i := 0; i < 20; i++ {
		printer.push("Print: " + strconv.Itoa(i))
	}
	printer.close()
}
