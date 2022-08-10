package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err1 := errors.New("error") // type: *errorString
	fmt.Printf("Type: %T, Value: %[1]v, errorString: %s\n", err1, err1.Error())

	err2 := errors.New("error")
	if errors.Is(err1, err2) {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同") // done， New 出来的，即使内容相同也不是同一个，不想等
	}

	_, err := os.Open("none.txt")

	// errors.As
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Println("Failed at path:", pathError.Path)
	} else {
		fmt.Println(err)
	}

	// errors.Unwrap
	e := errors.New("e")
	e1 := fmt.Errorf("e1: %w", e)
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := fmt.Errorf("e3: %w", e2)
	fmt.Printf("e: %#v, %[1]T\n", e)
	fmt.Printf("e1: %#v, %[1]T\n", e1)
	fmt.Printf("e2: %#v, %[1]T\n", e2)
	fmt.Printf("e3: %#v, %[1]T\n", e3)
	e10 := errors.Unwrap(e)
	e11 := errors.Unwrap(e1)
	e12 := errors.Unwrap(e2)
	e13 := errors.Unwrap(e3)
	fmt.Printf("Unwrap e10: %#v, %[1]T\n", e10)
	fmt.Printf("Unwrap e11: %#v, %[1]T\n", e11)
	fmt.Printf("Unwrap e12: %#v, %[1]T\n", e12)
	fmt.Printf("Unwrap e13: %#v, %[1]T\n", e13)
}
