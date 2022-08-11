package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type human struct {
	Name     string   `json:"name"`
	Sex      string   `json:"sex"`
	Age      int      `json:"age,string,omitempty"`
	Children []string `json:"children"`
	None     string   `json:"-"`
}

func main() {
	h10 := human{
		Name:     "qingbing",
		Sex:      "male",
		Age:      33,
		Children: []string{"nianyi", "chengyi"},
		None:     "Nothing",
	}

	// 编码到文件
	file10, err := os.OpenFile("human.out", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败2")
	} else {
		defer file10.Close()
		err = json.NewEncoder(file10).Encode(h10)
		if err != nil {
			fmt.Println("编码到文件失败")
		}
	}
	// 从文件中解码
	file20, err := os.OpenFile("human.out", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败2")
	} else {
		defer file20.Close()
		var h20 human
		err = json.NewDecoder(file20).Decode(&h20)
		if err != nil {
			fmt.Println("解码文件失败")
		} else {
			fmt.Printf("文件解码结果: %+v\n", h20)
		}
	}

	// 编码到文件，并格式化
	bs30, err := json.Marshal(h10)
	if err != err {
		fmt.Println("编码失败")
	}
	buffer30 := bytes.Buffer{}
	err = json.Indent(&buffer30, bs30, "", "\t")
	if err != nil {
		fmt.Println("编码到文件失败")
	}
	file30, err := os.OpenFile("human-indent.out", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败2")
	} else {
		defer file30.Close()
		buffer30.WriteTo(file30)
	}
}
