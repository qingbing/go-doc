package main

import "fmt"

func main() {
	// 数组定义
	a10 := [3]int{0, 1}
	// 切片定义
	s10 := []int{0, 1, 0}
	s11 := [...]int{0, 1, 0}

	fmt.Println("a10 = ", a10) // [0 1 0]
	fmt.Println("s10 = ", s10) // [0 1 0]
	fmt.Println("s11 = ", s11) // [0 1 0]

	a20 := [6]int{1, 2, 3, 4, 5, 6}
	// 切片从数组截取
	s20 := a20[2:4]
	fmt.Println("s20 = ", s20)           // [3 4]
	fmt.Println("len(s20) = ", len(s20)) // 1
	fmt.Println("cap(s20) = ", cap(s20)) // 4
}
