package main

import (
	"encoding/json"
	"fmt"
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
	// 普通编码
	s10, err := json.Marshal(h10)
	if err != nil {
		fmt.Println("编码失败 ==> ", err)
	} else {
		fmt.Println("编码结果 ==> ", string(s10))
	}
	// 格式化编码
	s11, err := json.MarshalIndent(h10, "", "\t")
	if err != nil {
		fmt.Println("编码失败 ==> ", err)
	} else {
		fmt.Println("编码结果 : \n", string(s11))
	}

	// 解码
	var h20 human
	err = json.Unmarshal(s11, &h20)
	if err != nil {
		fmt.Println("解码失败 ==> ", err)
	} else {
		fmt.Printf("解码结果 : %+v\n", h20)
	}
}
