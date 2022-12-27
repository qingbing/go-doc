package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	// 创建文件
	if f1, err := os.Create("file.out"); err!=nil {
		fmt.Println("Create Failed===>", err);
	}else{
		defer f1.Close()
	}
	// 只读方式打开文件
	if f2, err := os.Open("file.out"); err!= nil {
		fmt.Println("Open Failed===>", err);
	}else{
		defer f2.Close()
	}
	// 以定义权限方式打开文件
	if f3, err := os.OpenFile("file.out", os.O_RDWR, os.ModePerm); err != nil {
		fmt.Println("ReadFile Failed===>", err);
	}else{
		defer f3.Close()
		f3.WriteString("This is a test file\n")
		f3.WriteString("This is a test file\n")
		f3.WriteString("This is a test file\n")
		f3.WriteString("This is a test file")
		if n64, err := f3.Seek(5, io.SeekStart ); err != nil{
			fmt.Println("Seek fail ===> ", err)
		}else{
			if n, err := f3.WriteAt([]byte("a write goo"), n64); err != nil{
				fmt.Println("WriteAt fail ===> ", err)
			}else{
				fmt.Println("WriteAt Success: ", n)
			}
		}
	}
	/*
	读文件
	*/
	if f4, err := os.Open("file.out"); err != nil {
		fmt.Println("Open Failed===>", err);
	}else{
		defer f4.Close()
		reader := bufio.NewReader(f4)
		for{
			// 以换行符为分割符读取(读取一行)
			buffer, err := reader.ReadBytes('\n')
			f4.Write()
			if err !=nil && err == io.EOF {
				fmt.Print("Over: ", string(buffer))
				fmt.Println("Read Completed.")
				break
			} else if err != err {
				fmt.Println("ReadBytes failed ===> ", err)
				break
			}else{
				fmt.Print(string(buffer))
			}
		}

	}
}