package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// 文件查找位置
	fmt.Printf("SeekStart:%v\n", io.SeekStart)
	fmt.Printf("SeekCurrent:%v\n", io.SeekCurrent)
	fmt.Printf("SeekEnd:%v\n", io.SeekEnd)
	// Error
	fmt.Printf("Error - EOF:%v\n", io.EOF) // 文件末尾
	fmt.Printf("Error - ErrShortWrite:%v\n", io.ErrShortWrite)
	fmt.Printf("Error - ErrShortBuffer:%v\n", io.ErrShortBuffer)
	fmt.Printf("Error - ErrUnexpectedEOF:%v\n", io.ErrUnexpectedEOF)
	fmt.Printf("Error - ErrClosedPipe:%v\n", io.ErrClosedPipe)
	fmt.Printf("Error - ErrNoProgress:%v\n", io.ErrNoProgress)

	//func Copy(dst Writer, src Reader) (written int64, err error)

	if file, err := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal("打不开文件 test.txt", err)
	} else {
		defer file.Close()
		if n, err := io.Copy(os.Stdout, file); err != nil {
			log.Fatal("copy fail: ", err)
		} else {
			fmt.Printf("copy n: %d\n", n)
		}
	}

	// CopyBuffer
	fmt.Printf("===============\n")
	if file1, err := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal("打不开文件 test.txt", err)
	} else {
		defer file1.Close()
		var buf = make([]byte, 10)
		if n, err := io.CopyBuffer(os.Stdout, file1, buf); err != nil {
			log.Fatal("fail", err)
		} else {
			fmt.Println(n)
		}
	}

	// CopyN
	fmt.Printf("===============\n")
	if file2, err := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal("打不开文件 test.txt", err)
	} else {
		defer file2.Close()
		var totalN int64
		var readP int64 = 5
		for {
			n, err := io.CopyN(os.Stdout, file2, readP)
			if err == io.EOF {
				fmt.Println("读取完毕", totalN)
				break
			} else if err != nil {
				log.Fatal("fail", err)
			} else {
				totalN += n
			}
		}
	}
	// ReadAll
	fmt.Printf("===============\n")
	if file3, err := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal("打不开文件 test.txt", err)
	} else {
		defer file3.Close()
		if bs, err := io.ReadAll(file3); err != nil {
			log.Fatal("fail", err)
		} else {
			fmt.Println(bs)
		}
	}

	// func WriteString(w Writer, s string) (n int, err error)
	fmt.Printf("===============\n")
	if file, err := os.OpenFile("write.out", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm); err != nil {
		log.Fatal("打不开文件 write.out", err)
	} else {
		if n, err := io.WriteString(file, "just test\n"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(n)
		}
	}
	// func ReadFull(r Reader, buf []byte) (n int, err error):
	fmt.Printf("===============\n")
	if file, err := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm); err != nil {
		log.Fatal("打不开文件 test.txt", err)
	} else {
		defer file.Close()
		var buf = make([]byte, 100)
		n, err := io.ReadFull(file, buf)
		if err != nil {
			log.Fatal("fail", err)
		} else {
			fmt.Println(n, string(buf))
		}
	}

	// func Pipe() (*PipeReader, *PipeWriter)
	fmt.Printf("===============\n")
	r, w := io.Pipe() // 返回一个 PipeReader 和 一个 PipeWriter， 任何一个被操作，都会进行阻塞，所以，需要必须开协程
	lock := make(chan struct{})
	go func() {
		w.Write([]byte("Hello Pipe"))
		lock <- struct{}{}
	}()
	buf := make([]byte, 100)
	r.Read(buf)
	fmt.Println(string(buf))
	<-lock

	// func LimitReader(r Reader, n int64) Reader
	fmt.Printf("====LimitReader===========\n")
	file, _ := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm)
	r01 := io.LimitReader(file, 10)
	io.Copy(os.Stdout, r01)
	r01 = io.LimitReader(file, 10)
	io.Copy(os.Stdout, r01)

	// func TeeReader(r Reader, w Writer) Reader
	fmt.Printf("\n====TeeReader===========\n")
	var buf01 bytes.Buffer
	r02 := io.TeeReader(file, &buf01)
	_, err := io.Copy(os.Stdout, r02)
	if err != nil {
		log.Fatal("fail", err)
	}

	// func MultiReader(readers ...Reader) Reader
	fmt.Printf("====MultiReader===========\n")
	//file, _ := os.OpenFile("test.txt", os.O_RDONLY, os.ModePerm)
	reader21 := strings.NewReader("hello")
	reader22 := strings.NewReader("world")
	r03 := io.MultiReader(file, reader21, reader22)
	io.Copy(os.Stdout, r03)

	// func MultiWriter(writers ...Writer) Writer
	fmt.Printf("====MultiWriter===========\n")
	file01, _ := os.OpenFile("m1.out", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file02, _ := os.OpenFile("m2.out", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	r04 := io.MultiWriter(file01, file02)
	r04.Write([]byte("ok"))

	// func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
	fmt.Printf("\n===NewSectionReader============\n")
	reader := strings.NewReader("Geeks\n")
	sectionReader := io.NewSectionReader(reader, 3, 5)
	io.Copy(os.Stdout, sectionReader)

	//fmt.Printf("\n===============\n")
	// func NopCloser(r Reader) ReadCloser
}
