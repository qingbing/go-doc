package main

import "fmt"

/*
slice-去除空字符串
*/
// 使用 append 新开切片进行去除空字符串处理
func noEmpty1(data []string) (out []string) {
	for _, str := range data {
		if str == "" {
			continue
		}
		out = append(out, str)
	}
	return
}

// 不使用 append 进行去除空字符串处理
func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if "" == str {
			continue
		}
		data[i] = str
		i++
	}
	return data[:i]
}

func main() {
	data := []string{"red", "", "blue", "yellow", "", "pink"}
	fmt.Println("data:", data, ": len=", len(data), "; cap=", cap(data))

	afterData1 := noEmpty1(data)
	fmt.Println("afterData1:", afterData1, "; len=", len(afterData1), "; cap=", cap(afterData1))

	afterData2 := noEmpty2(data)
	fmt.Println("afterData2:", afterData2, "; len=", len(afterData2), "; cap=", cap(afterData2))
}
