package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	p, err := os.FindProcess(os.Getpid())
	fmt.Println(p)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Signal(os.Interrupt); err != nil {
		log.Fatal(err)
	}

	select {
	case <-time.After(time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context canceled"
		stop()                 // stop receiving signal notifications as soon as possible.
	}

	time.Sleep(time.Second * 10)
}
