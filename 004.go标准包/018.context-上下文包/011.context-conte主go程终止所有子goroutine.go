package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func doThing(ctx context.Context, name string) {
	if name == "sleeping" {
		go doSubThing(ctx, "subthing")
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "over")
			return
		case <-time.After(time.Second * time.Duration(rand.Intn(3)+1)):
			fmt.Println(name, time.Now())
		}
	}
}

func doSubThing(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "over")
			return
		case <-time.After(time.Second * time.Duration(rand.Intn(3)+1)):
			fmt.Println(name, time.Now())
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go doThing(ctx, "eating")
	go doThing(ctx, "sleeping")
	go doThing(context.WithValue(ctx, "name", "xxx"), "shoping")

	time.Sleep(time.Second * 7)
	cancel()
	time.Sleep(time.Second * 2)
}
