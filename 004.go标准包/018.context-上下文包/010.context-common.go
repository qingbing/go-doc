package main

import (
	"context"
	"fmt"
	"time"
)

func valueCancel(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	fmt.Println("睡觉，传递内容: ", ctx.Value("id"))
	time.Sleep(time.Second * 4)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		fmt.Println(time.Now())
		time.Sleep(time.Second * 1)
	}()
	<-ctx.Done()
	fmt.Println(ctx.Err())
	fmt.Println(time.Now())

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	<-ctx.Done()
	fmt.Println(ctx.Err())
	fmt.Println(time.Now())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	<-ctx.Done()
	fmt.Println(ctx.Err())
	fmt.Println(time.Now())

	ctx = context.WithValue(context.Background(), "id", "name")
	ctx, cancel = context.WithCancel(ctx)
	go valueCancel(ctx, cancel)
	<-ctx.Done()
	fmt.Println(ctx.Err())
	fmt.Println(time.Now())
}
