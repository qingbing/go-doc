package main

import (
	"fmt"
	"path/filepath"
)

func realpath() {
	fmt.Println(filepath.Abs("."))
}

func main() {
	realpath()

}
