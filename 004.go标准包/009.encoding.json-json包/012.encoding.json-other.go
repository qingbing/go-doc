package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	buffer10 := bytes.Buffer{}
	s10 := "<qingbing>&<dd>ending"
	// <, >, &, u+2028, u+2029 转化成 unicode 编码后写入 buffer
	json.HTMLEscape(&buffer10, []byte(s10))
	// 把json数据添加到buffer中去
	bs, _ := json.Marshal(map[string]int{"qing": 11})
	err := json.Compact(&buffer10, []byte("11ok"))
	err = json.Compact(&buffer10, bs)
	if err != nil {
		fmt.Println(err)
	}
	buffer10.WriteTo(os.Stdout)
	fmt.Println()
}
