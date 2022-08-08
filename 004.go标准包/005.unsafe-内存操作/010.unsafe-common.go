package main

import (
	"fmt"
	"unsafe"
)

type ts struct {
	i   byte
	i32 int32
	b   bool
}
type human struct {
	name string
	sex  string
	age  uint8
}

func main() {
	fmt.Printf("int, Sizeof: %d; Alignof: %d\n", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)))          // 1, 1
	fmt.Printf("int, Sizeof: %d; Alignof: %d\n", unsafe.Sizeof(int(0)), unsafe.Alignof(int(0)))            // 8, 8
	fmt.Printf("string, Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(string("")), unsafe.Alignof(string(""))) // 16, 8
	sli1 := []string{"1"}
	fmt.Printf("slice, Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(sli1), unsafe.Alignof(sli1)) // 24, 8

	sli2 := []int8{1, 2, 3, 4, 5, 6}
	fmt.Printf("slice, Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(sli2), unsafe.Alignof(sli2)) // 12, 4

	ts1 := ts{}
	fmt.Printf("ts-struct, Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(ts1), unsafe.Alignof(ts1)) // 40, 8
	h := human{"qingbing", "male", 11}
	fmt.Printf("human-struct, Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(h), unsafe.Alignof(h)) // 40, 8

	// 获取并修改
	fmt.Println("修改前 ===> ", h)
	hPtr := unsafe.Pointer(&h)
	namePtr := (*string)(hPtr) // (*string)(unsafe.Pointer(uintptr(hPtr) + unsafe.Offsetof(h.name))) // 默认第一个字段偏移量为 0
	sexPtr := (*string)(unsafe.Pointer(uintptr(hPtr) + unsafe.Offsetof(h.sex)))
	agePtr := (*uint)(unsafe.Pointer(uintptr(hPtr) + unsafe.Offsetof(h.age)))
	fmt.Println(*namePtr, *sexPtr, *agePtr)
	*namePtr = "yongjing"
	*sexPtr = "female"
	*agePtr = 22
	fmt.Println("修改后 ===> ", h)
}
