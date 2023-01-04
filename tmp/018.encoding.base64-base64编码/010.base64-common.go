package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}
func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func main() {
	str := "hello world"
	bStr01 := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(bStr01)
	if s01, err := base64.StdEncoding.DecodeString(bStr01); err != nil {
		fmt.Println("解析失败")
	} else {
		fmt.Println(string(s01))
	}

	var es02 []byte
	base64.StdEncoding.Encode([]byte(str), es02)
	fmt.Println(string(es02))

	//debyte := base64Encode([]byte(hello))
	//// decode
	//enbyte, err := base64Decode(debyte)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if hello != string(enbyte) {
	//	fmt.Println("hello is not equal to enbyte")
	//}
	//fmt.Println(string(enbyte))
}
