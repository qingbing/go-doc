package main

import (
	"flag"
	"fmt"
)

/**
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
	- errorHandling
		- flag.ContinueOnError: 发生错误后继续解析，CommandLine就是使用这个选项
		- flag.ExitOnError: 出错时调用os.Exit(2)退出程序
		- flag.PanicOnError: 出错时产生 panic
*/
// -iFlag=11 -bFlag=true -sFlag=ok
func main() {
	args := []string{"-iFlag", "1", "-sFlag", "test", "-bFlag", "false"} // 需要解析的参数必须

	var iFlag int
	var bFlag bool
	var sFlag string

	fs := flag.NewFlagSet("myFlagSet", flag.PanicOnError)
	fs.IntVar(&iFlag, "iFlag", 0, "int flag value")
	fs.BoolVar(&bFlag, "bFlag", false, "bool flag value")
	fs.StringVar(&sFlag, "sFlag", "default", "string flag value")

	fs.Parse(args) // 自定义 flag 解析时，需要显示传入字符串切片作为参数

	fmt.Println("int flag:", iFlag)
	fmt.Println("bool flag:", bFlag)
	fmt.Println("string flag:", sFlag)
}
