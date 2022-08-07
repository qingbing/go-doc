package main

// 多个选项指定到同一变量(短选项实现)
import (
	"flag"
	"fmt"
)

var (
	level string
)

func init() {
	var (
		levelDefault = ""
		levelDesc    = "level description"
	)
	flag.StringVar(&level, "ll", levelDefault, levelDesc)
	flag.StringVar(&level, "sl", levelDefault, levelDesc)
}

// -ll=debug
// 等同
// -sl=debug
func main() {
	flag.Parse()
	fmt.Println("level:", level)
}
