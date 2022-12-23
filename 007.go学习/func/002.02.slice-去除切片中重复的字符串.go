package main

import "fmt"

/*
slice-去除切片中重复的字符串
*/
func noRepeat(data []string) (out []string) {

	for _, word := range data {
		i := 0
		len := len(out)
		for ; i < len; i++ {
			if word == out[i] {
				break
			}
		}
		if i == len {
			out = append(out, word)
		}
	}
	return
}

func main() {
	data := []string{"red", "blue", "yellow", "yellow", "pink", "pink", "blue"}
	fmt.Println("data=", data, "; len=", len(data), "; cap=", cap(data))

	data1 := noRepeat(data)
	fmt.Println("data1=", data1, "; len=", len(data1), "; cap=", cap(data1))
}
