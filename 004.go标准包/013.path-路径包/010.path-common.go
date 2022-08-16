package main

import (
	"fmt"
	"path"
)

func main() {
	var p01 string
	p01 = "t1/t2/t3.txt"
	fmt.Printf("(%s) base: (%s)\n", p01, path.Base(p01)) // t3.txt
	p01 = "t1/./../t2/../t3.txt"
	fmt.Printf("(%s) Clean: (%s)\n", p01, path.Clean(p01)) // t3.txt
	p01 = "t1/t2/t3.txt"
	fmt.Printf("(%s) Dir: (%s)\n", p01, path.Dir(p01)) // t1/t2
	p01 = "t1/t2/t3.txt"
	fmt.Printf("(%s) Ext: (%s)\n", p01, path.Ext(p01)) // .txt
	p01 = "t1/t2/t3.txt"
	fmt.Printf("(%s) IsAbs: (%t)\n", p01, path.IsAbs(p01)) // false
	p01 = "/tmp"
	fmt.Printf("(%s) IsAbs: (%t)\n", p01, path.IsAbs(p01))   // true
	fmt.Printf("%s\n", path.Join("bb", "//c", "dd//", "ee")) // bb/c/dd/ee
	p01 = "t1/t2/t3.txt"
	p02, filename := path.Split(p01)
	fmt.Printf("(%s) dir: (%s); file: (%s)\n", p01, p02, filename) // (t1/t2/t3.txt) dir: (t1/t2/); file: (t3.txt)

	p01 = "t1/t2/t3.txt"
	pattern := "t1/*/5?.txt"
	isMatch, err := path.Match(pattern, p01)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("(%s) match (%s): (%t)\n", p01, pattern, isMatch)
	}
}
