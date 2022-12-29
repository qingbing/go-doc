package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
file-按行读取文件，并且将行内容用回调函数进行处理
*/
func ScanFileByLine(filename string, callback func(sentence string) (isContinue bool, err error)) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	bs := []byte{}
	isContinue := true
	for {
		bs, err = reader.ReadBytes('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			return
		}
		isContinue, err = callback(string(bs))
		if err != nil {
			return
		} else if isContinue == false {
			return nil
		}
	}
}

/*
去除字符串中的标点符号
*/
func TrimDot(str string) (res string) {
	reg, _ := regexp.Compile("[\"\n',;:*.!?|\\\\\\[\\]&=>()/{}]")
	return reg.ReplaceAllString(str, " ")
}

func main() {
	filename := "test.out"
	countInt := 0
	countWord := "Person"
	err := ScanFileByLine(filename, func(sentence string) (bool, error) {
		sentence = TrimDot(sentence)
		words := strings.Fields(sentence)
		for _, word := range words {
			if word == countWord {
				countInt++
			}
		}
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("File(%s) has word(%s): %d\n", filename, countWord, countInt)
	}
}
