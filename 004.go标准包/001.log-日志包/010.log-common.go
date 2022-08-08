package main

import (
	"fmt"
	"log"
	"os"
)

type user struct {
	name string
	age  int
}

func main() {
	person := user{
		name: "qingbing",
		age:  18,
	}
	fmt.Println("Current Flags ===> ", log.Flags())
	log.Println("Println", person)
	log.SetPrefix("TestPrefix: ")
	log.Printf("Printf, Name: %s, Age: %d", person.name, person.age)
	//log.SetFlags(log.Ldate | log.Llongfile | log.LstdFlags)
	fmt.Println("Prefix ===> ", log.Prefix())
	fmt.Println("==============")
	// 设置自定义的logger，打印到自定义的 writer 上
	mywriter, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal("Open file fail.")
	}
	logger := log.New(mywriter, "My Log: ", log.Ldate|log.Ltime)
	logger.Println("My log first")
	logger.Println("My log second")
}
