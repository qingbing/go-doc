package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	p1 := "./.."
	realPath, err := filepath.Abs(p1)
	if err == nil {
		fmt.Printf("(%s) Abslote path is (%s)\n", p1, realPath)
	} else {
		fmt.Println(err)
	}

	fmt.Printf("(%s) IsAbs: %t\n", p1, filepath.IsAbs(p1))
	fmt.Printf("(%s) IsAbs: %t\n", realPath, filepath.IsAbs(realPath))
	fmt.Printf("(%s) Base: %s\n", p1, filepath.Base(p1))
	fmt.Printf("(%s) Base: %s\n", realPath, filepath.Base(realPath))
	fmt.Printf("(%s) Dir: %s\n", p1, filepath.Dir(p1))
	fmt.Printf("(%s) Dir: %s\n", realPath, filepath.Dir(realPath))

	// ext
	file := "tmp/tt.txt"
	fmt.Printf("(%s) Ext: %s\n", realPath, filepath.Ext(realPath))
	fmt.Printf("(%s) Ext: %s\n", file, filepath.Ext(file))
	// rel
	rp1 := realPath + "/tmp/xx/eex.txt"
	relativePath, err := filepath.Rel(realPath, rp1)
	if err == nil {
		fmt.Printf("(%s, %s) Rel: %s\n", realPath, rp1, relativePath)
	} else {
		fmt.Println(err)
	}

	fmt.Println(filepath.Join("tmp", "xx/txt.txt"))
	fmt.Println(filepath.Split("tmp/xx/txt.txt"))
	fmt.Println(filepath.SplitList("/home/qingbing/code:/usr/local/go:"))
	fmt.Println(filepath.Clean("/home/qingbing/code/../../"))
	fmt.Println(filepath.ToSlash("/home/wohu/GoCode/src"))
	fmt.Println(filepath.VolumeName("/home/wohu/GoCode/src"))

	fmt.Println("=== WalkDir 功能同 Walk 效率高于 wall，避免了在每个访问的文件或目录调用 os.Lstat ===")
	filepath.WalkDir(realPath, func(path string, info fs.DirEntry, err error) error {
		fmt.Println(path, "====>", info.Name())
		return nil
	})
	//fmt.Println("=== walk ===")
	//filepath.Walk(realPath, func(path string, info fs.FileInfo, err error) error {
	//	fmt.Println(path, "====>", info.Name())
	//	return nil
	//})

}
