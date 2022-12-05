package main

import "fmt"

// func test(m int) *string {
//     b := 5
//     fmt.Println(m + b)
//     var p *string = new(string)
//     *p = "hello"
//     fmt.Printf("%q\n", *p) // 打印 go 语言格式的字符串
//     return p
// }
// func main() {
//     a := 100
//     fmt.Println(a)
//     fmt.Println(a)
//     p := test(10)
//     fmt.Printf("%v", *p)
// }

func swap(x, y *int) {
	*x, *y = *y, *x
	fmt.Printf("a: %d, b: %d\n", *x, *y)
}

func main() {
	a, b := 10, 20
	fmt.Printf("a: %p, b: %p\n", &a, &b)
	swap(&a, &b)
	fmt.Printf("a: %p, b: %p\n", &a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
}
