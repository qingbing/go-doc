package main

import (
	"flag"
	"fmt"
)

var (
	i int
	b bool
	s string
)

func init() {
	// 变量地址直接引用
	flag.IntVar(&i, "i1", 0, "输入 int")
	flag.BoolVar(&b, "b1", false, "输入 boo")
	flag.StringVar(&s, "s1", "", "输入 string")
}

// -i1=2 -- xx -b1 -s1=sss
func main() {
	// 选项解析
	flag.Parse()
	// 输出解析
	fmt.Println("int", i)
	fmt.Println("bool", b)
	fmt.Println("string", s)

	// 终止解析后的参数
	// flag.Args() 未解析的参数切片
	// flag.NArg() 未解析的参数个数
	// flag.arg(int i) 第 i 个参数，i 从 0 开始
	//fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("第 %d 个参数: %v\n", i, flag.Arg(i))
	}
}
