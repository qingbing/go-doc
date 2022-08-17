package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	filename := "test.txt"
	path := "."
	// func ReadAll(r io.Reader) ([]byte, error)
	fmt.Println("====== ReadAll ==============")
	f1, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败", err)
	} else {
		defer f1.Close()
		bs, err := ioutil.ReadAll(f1)
		if err != nil {
			fmt.Println("文件读取失败", err)
		} else {
			fmt.Println(string(bs))
		}
	}
	// func ReadDir(dirname string) ([]fs.FileInfo, error)
	fmt.Println("====== ReadDir ==============")
	if fs, err := ioutil.ReadDir(path); err != nil {
		fmt.Println("读取目录失败", err)
	} else {
		for _, fileInfo := range fs {
			fmt.Printf("Name: %s; IsDir: %t; Size: %d\n", fileInfo.Name(), fileInfo.IsDir(), fileInfo.Size())
		}
	}
	// func ReadFile(filename string) ([]byte, error)
	fmt.Println("====== ReadFile ==============")
	if bs, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("文件读取失败", err)
	} else {
		fmt.Println(string(bs))
		// func WriteFile(filename string, data []byte, perm fs.FileMode) error
		fmt.Println("====== WriteFile ==============")
		nbs := bytes.NewBuffer(bs)
		nbs.WriteString("Add ")
		if err := ioutil.WriteFile("test.out", nbs.Bytes(), os.ModePerm); err != nil {
			fmt.Println("文件写入失败", err)
		} else {
			fmt.Println("文件写入成功", err)
		}
	}

	// func TempDir(dir, pattern string) (name string, err error)
	fmt.Println("====== TempDir ==============")
	if dir, err := ioutil.TempDir(".", "test"); err != nil {
		fmt.Println("创建临时目录失败", err)
	} else {
		defer os.Remove(dir)
		fmt.Println(dir)
		// func TempFile(dir, pattern string) (f *os.File, err error)
		fmt.Println("====== TempFile ==============")
		if f2, err := ioutil.TempFile(dir, "test"); err != nil {
			fmt.Println("创建临时目录失败", err)
		} else {
			defer os.Remove(f2.Name())
			defer f2.Close()
			fmt.Println(f2.Name())
			f2.WriteString("Just Test")
		}
	}

	// func NopCloser(r io.Reader) io.ReadCloser
	fmt.Println("====== NopCloser ==============")
	httpReq := http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("hello"))),
	} // 一般从页面传递进来
	data, err := ioutil.ReadAll(httpReq.Body)
	// coding ... 获取到了 body 数据，操作数据
	fmt.Println(string(data))
	// 将数据装填会 body
	httpReq.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	// coding ... 至此就可以和之前一样使用 request 及里面的数据了
	data, err = ioutil.ReadAll(httpReq.Body)
	fmt.Println(string(data))
}
