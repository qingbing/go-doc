package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

// dp.ReadDir(-1) : 读取所有文件内容
/*
file-遍历目录
*/
func ScanDir(dir string, callback func(file fs.DirEntry, parentDir string) (isContinue bool, err error)) (err error) {
	finfo, err := os.Stat(dir)
	if err != nil {
		return
	} else if !finfo.IsDir() {
		return fmt.Errorf("%s is not a directory", dir)
	}
	// 打开目录读取句柄
	dp, err := os.OpenFile(dir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return
	}
	defer dp.Close()

	files := []fs.DirEntry{}
	var isContinue bool
	for {
		files, err = dp.ReadDir(50)
		if err != nil && err != io.EOF {
			return
		}
		if len(files) == 0 {
			// 读取完毕
			return nil
		}
		for _, file := range files {
			if isContinue, err = callback(file, dir); err != nil {
				return
			} else if false == isContinue {
				return
			}
		}
	}
}

func main() {
	dir := "."
	goFile := []string{}
	err := ScanDir(dir, func(file fs.DirEntry, dir string) (bool, error) {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			goFile = append(goFile, dir+"/"+file.Name())
		}
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
	} else {
		for _, filename := range goFile {
			fmt.Printf("%s\n", filename)
		}
	}
}
