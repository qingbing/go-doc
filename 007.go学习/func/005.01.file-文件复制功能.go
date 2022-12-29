package main

import (
	"fmt"
	"io"
	"os"
)

/*
file-文件拷贝
*/
func CopyFile(src, dist string) (copySize int, err error) {
	// 打开源文件
	sFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer sFile.Close()

	// 打开写入文件
	dFile, err := os.Create(dist)
	if err != nil {
		return
	}
	defer dFile.Close()

	// 创建缓冲区，边读边写
	buf := make([]byte, 4096) // 由于虚拟内存的最小单位是 page(默认是4096)，所以，设置成 4096
	readSize := 0
	for {
		readSize, err = sFile.Read(buf)
		if err != nil && err != io.EOF { // 读取错误
			return
		}
		if readSize == 0 { // 读完毕
			err = nil
			return
		}
		copySize += readSize
		// 写入文件
		dFile.Write(buf[:readSize])
	}
}

func main() {
	filename := "xx.tar.gz"
	distFile := "go.tar.gz"
	if size, err := CopyFile(filename, distFile); err != nil {
		fmt.Println("Copy file failed.", err)
	} else {
		fmt.Println("Copy file success: ", size)
	}
}
