package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	bs01 := []byte("I am ok")
	bs02 := []byte("I am oK")
	fmt.Println(string(bytes.ToLower(bs01)))
	fmt.Println(string(bytes.ToUpper(bs01)))
	fmt.Println(len(bs01))
	fmt.Println(string(bs01))
	fmt.Println(bytes.Compare(bs01, bs02))
	fmt.Println(bytes.Equal(bs01, bs02))     // 比较两个字符串是否相等, 区分大小写
	fmt.Println(bytes.EqualFold(bs01, bs02)) // 比较两个字符串是否相等, 不区分大小写

	fmt.Printf("bytes.MinRead: %d\n", bytes.MinRead) // Buffer.ReadFrom传递给Read调用的最小切片大小

	fmt.Println(bytes.Count([]byte("hello"), []byte("l"))) // count("l")
	fmt.Println(bytes.Count([]byte("hello"), []byte("")))  // len(s)+1
	fmt.Println(bytes.Runes([]byte("hello中")))             // 返回 unicode 码表示

	fmt.Println(strings.ToValidUTF8("Geeks\xc5Geeks", "For")) // 替换无效的 unicode

	// buffer 操作
	//buf := bytes.Buffer{}
	//buf := bytes.NewBuffer([]byte{})
	buf := bytes.NewBufferString("")
	buf.Write([]byte("hello"))
	buf.WriteString(" ")
	buf.WriteString("world")
	buf.WriteString("\n")
	buf.WriteTo(os.Stdout)
}
