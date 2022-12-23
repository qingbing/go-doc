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

	// 创建切片
	fmt.Println("创建切片")
	s31 := []int{1, 2, 3}
	fmt.Println("s31, len=", len(s31), "cap=", cap(s31))
	s32 := make([]int, 2)
	fmt.Println("s32, len=", len(s32), "cap=", cap(s32))
	s33 := make([]int, 2, 3)
	fmt.Println("s33, len=", len(s33), "cap=", cap(s33))

	// 切片追加 append
	s41 := []int{0, 1, 2}
	s41 = append(s41, 4)
	fmt.Println("s41, len=", len(s41), "cap=", cap(s41))

	// 切片拷贝 copy
	s51 := []int{0, 1, 2}
	s52 := []int{3, 4}
	i51 := copy(s51, s52)
	fmt.Println("i51:", i51) //2
	fmt.Println("s52:", s52) // [3, 4]
	fmt.Println("s51:", s51) // [3, 4, 2]

	s61 := []int{0, 1, 2}
	s62 := []int{3, 4}
	i61 := copy(s62, s61)
	fmt.Println("i61:", i61) // 3
	fmt.Println("s61:", s61) // [0, 1, 2]
	fmt.Println("s62:", s62) // [0, 1]
}
