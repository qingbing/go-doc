package main

import "fmt"

// struct 类型定义
type Person struct{
    name string
    age int
    sex int
}

func main(){
	// 初始化
	p1 := Person{"qingbing", 33, 1}
	fmt.Println("p1 ===>", p1)

	p2 := Person{
		name: "yong",
		sex:2,
	}
	fmt.Println("p2 ===>", p2)

	var p3 Person
	p3.name = "test"
	fmt.Println("p3 ===>", p3)

	// 结构体的比较
	p4 := Person{"test", 1, 1}
	p5 := Person{"test", 1, 1}
	p6 := Person{"test1", 1, 1}
	fmt.Printf("p4 == p5: %t\n", p4 == p5)
	fmt.Printf("p4 == p6: %t\n", p4 == p6)
}