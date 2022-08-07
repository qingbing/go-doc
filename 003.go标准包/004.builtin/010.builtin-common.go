package main

import "fmt"

func main() {
	sli1 := make([]int, 1, 3)
	fmt.Printf("Length: %d, Cap: %d, Slice: %v\n", len(sli1), cap(sli1), sli1)
	sli1 = append(sli1, 1)
	sli1 = append(sli1, 2)
	sli1 = append(sli1, 3)
	fmt.Printf("Length: %d, Cap: %d, Slice: %v\n", len(sli1), cap(sli1), sli1)

	str1 := new(string)
	fmt.Printf("string: %s, len: %d, type: %[1]T\n", *str1, len(*str1))

	sli2 := new([]string)
	*sli2 = append(*sli2, "qing")
	fmt.Printf("Length: %d, Cap: %d, Slice: %v\n", len(*sli2), cap(*sli2), *sli2)

	sli3 := make([]string, 0)
	count := copy(*sli2, sli3) // 注意 copy 时不能时指针
	fmt.Printf("Copy leng: %d, Cap: %d, Slice: %v\n", count, cap(sli3), sli3)
}
