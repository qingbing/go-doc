
# main参数的接受

- 参考文件
  - [flag-命令行解析包](./../../004.go标准包/002.flag-命令行解析/001.flag-命令行解析包.md)


## 1. 代码

```go
/*
获取帮助信息: ./demo -h
参数指定:./demo -s=str -b=true -i=10 -short=sh -ps=ptr
*/
package main

import (
	"flag"
	"fmt"
)

var (
	sVal string
	bVal bool
	iVal int

	lsVal string

	ptrSVal *string
)

func init() {
	/* 定义接收参数 */
	// 普通接收参数
	flag.StringVar(&sVal, "s", "default", "接收string")
	flag.BoolVar(&bVal, "b", false, "接收bool")
	flag.IntVar(&iVal, "i", 0, "接收int")

	// 长短标签(多标签控制变量)
	flag.StringVar(&lsVal, "short", "", "短标签")
	flag.StringVar(&lsVal, "long", "", "长标签")

	// 指针标签定义方式
	ptrSVal = flag.String("ps", "ptr", "指针标签")
	/* 解析参数 */
	flag.Parse()
}

func main() {
	fmt.Printf("String Value: %s\n", sVal)
	fmt.Printf("Bool Value: %t\n", bVal)
	fmt.Printf("Int Value: %d\n", iVal)
	fmt.Printf("长短标签: %s\n", lsVal)
	fmt.Printf("指针标签: %s\n", *ptrSVal)
}
```