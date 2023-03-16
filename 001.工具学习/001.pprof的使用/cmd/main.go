package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func subTest1() {
	fmt.Println("subTest1 -- start")
	for i := 0; i < 100000; i++ {
		fmt.Print(1)
	}
	fmt.Println("subTest1 -- end")
}
func subTest2() {
	fmt.Println("subTest2 -- start")
	val := 0
	for i := 0; i < 1000000; i++ {
		val = i
		fmt.Print(val)
	}
	fmt.Println("subTest2 -- end")
}
func subTest3() {
	fmt.Println("subTest3 -- start")
	val := 0
	for i := 0; i < 1000000; i++ {
		val = i
		fmt.Print(val)
	}
	fmt.Println("subTest3 -- end")
}

func mainFun() {
	fmt.Println("mainFun -- start")
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Print(0)
		}
	}()
	go subTest1()
	go subTest2()
	subTest3()
	fmt.Println("mainFun -- end")
}

func main() {
	// 记录 cpu profiling
	cpuFile, err := os.Create("cpu.log")
	if err != nil {
		log.Fatalln("打开文件失败")
	}
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)

	defer pprof.StopCPUProfile()
	// 记录 mem profiling
	memFile, err := os.Create("mem.log")
	if err != nil {
		log.Fatalln("打开文件失败")
	}
	defer memFile.Close()
	pprof.WriteHeapProfile(memFile)

	mainFun()
}
