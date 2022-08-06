package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
	timeFlag   time.Duration

	intFlag1    *int
	boolFlag1   *bool
	stringFlag1 *string
)

func init() {
	// 变量地址直接引用
	flag.IntVar(&intFlag, "i1", 0, "输入 int")
	flag.BoolVar(&boolFlag, "b1", false, "输入 boo")
	flag.StringVar(&stringFlag, "s1", "", "输入 string")
	flag.DurationVar(&timeFlag, "t1", time.Second*1, "输入 time")

	// 返回变量指针
	intFlag1 = flag.Int("i2", 1, "请输入 int")
	boolFlag1 = flag.Bool("b2", false, "请输入 bool")
	stringFlag1 = flag.String("s2", "", "请输入 string")
}

func main() {
	// 选项解析
	flag.Parse()
	// 输出解析
	fmt.Println("intFlag", intFlag)
	fmt.Println("boolFlag", boolFlag)
	fmt.Println("stringFlag", stringFlag)
	fmt.Println("timeFlag", timeFlag)

	fmt.Println("intFlag1", intFlag1)
	fmt.Println("boolFlag1", boolFlag1)
	fmt.Println("stringFlag1", stringFlag1)
}
