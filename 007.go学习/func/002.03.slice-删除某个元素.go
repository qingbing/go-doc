package main

import "fmt"

/*
slice-删除某个元素
*/
// 使用 copy 方法
func removeSlice1(data []string, word string) []string {
	for i, w := range data {
		if w == word {
			copy(data[i:], data[i+1:])
			return removeSlice1(data[:len(data)-1], word)
		}
	}
	return data
}

// 推荐方式
func removeSlice2(data []string, word string) []string {
	c := 0
	for i := 0; i < len(data); i++ {
		if word == data[i] {
			continue
		}
		data[c] = data[i]
		c++
	}
	return data[:c]
}

func main() {
	data1 := []string{"red", "blue", "yellow", "blue", "pink"}
	afterData1 := removeSlice1(data1, "blue")
	fmt.Println("data1: => ", data1)
	fmt.Println("afterData1: => ", afterData1)

	data2 := []string{"red", "blue", "yellow", "blue", "pink"}
	afterData2 := removeSlice2(data2, "blue")
	fmt.Println("data2: => ", data2)
	fmt.Println("afterData2: => ", afterData2)
}
