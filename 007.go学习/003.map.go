package main

import (
	"fmt"
)

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
	fmt.Println("m4=", m4, "; len=", len(m4))

	var m5 map[string]string = map[string]string{
		"name" : "qingbing",
		"sex":"nan",
	}
	fmt.Println("m5=", m5, "; len=", len(m5))

	var m6 map[string]string = map[string]string{
		"name" : "qingbing",
		"sex":"nan",
	}
	for k, v := range m6{
		fmt.Println("key:", k, "; value:", v)
	}
	for k := range m6{
		fmt.Println("key:", k)
	}
	for _,v := range m6{
		fmt.Println("value:", v)
	}
	
	if v, ok := m6["xx"]; ok{
		fmt.Println("存在 key:", "xx", "; 值为=", v)
	}else{
		fmt.Println("不存在 key:", "xx")
	}

	// map 的删除使用
	var m7 map[string]string = map[string]string{
		"name" : "qingbing",
		"sex":"nan",
	}
	delete(m7, "sex")
	delete(m7, "age")
	fmt.Println("m7=", m7, "; len=", len(m7))
}
