package main

import "fmt"

func main() {
	// 创建 map
	var m1 map[string]string
	if m1 == nil {
		fmt.Println("m1 是 nil 的，声明时没有空间，不能存储数据")
	}

	m2 := map[string]string{}
	m2["name"] = "qingbing"
	fmt.Println("m2=", m2, "; len=", len(m2))

	m3 := make(map[string]string)
	m3["name"] = "qingbing"
	fmt.Println("m3=", m3, "; len=", len(m3))

	m4 := make(map[string]string, 5)
	m4["name"] = "qingbing"
	fmt.Println("m4=", m4, "; len=", cap(m4))
}
