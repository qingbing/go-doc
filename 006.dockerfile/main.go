package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Millisecond * 300)
		fmt.Printf("Line:%-5d; Hello World\n", i)
	}
}
