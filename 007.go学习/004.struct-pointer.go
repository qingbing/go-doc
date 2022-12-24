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
	ptr1 := &Person{"qingbing", 33, 1}
	fmt.Println("ptr1 ===>", ptr1)

	ptr2 := &Person{
		name: "yong",
		sex:2,
	}
	fmt.Println("ptr2 ===>", ptr2)


	p3 := Person{age: 1}
	var ptr3 *Person = &p3
	ptr3.name = "test"
	fmt.Println("ptr3 ===>", ptr3)

	ptr4 := new(Person)
	ptr4.name = "test4"
	fmt.Println("ptr4 ===>", ptr4)
	fmt.Println("ptr4 ===>", ptr4)
}