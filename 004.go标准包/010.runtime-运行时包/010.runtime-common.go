package main

import (
	"fmt"
	"runtime"
	"time"
)

type human struct {
	name string
}

func main() {
	runtime.GOMAXPROCS(2) // 设置最大可同时使用的 cpu 核数
	//go func() {
	//	fmt.Println("before Gosched")
	//	runtime.Gosched() // 手动退出当前 goroutine，让出 CPU， 后面还是会执行
	//	fmt.Println("after Gosched")
	//}()
	//go func() {
	//	defer func() {
	//		fmt.Println("defer goexit")
	//	}()
	//	runtime.Goexit()            // 退出当前 goroutine，后面的代码不执行，但是已注册的 defer 函数会执行
	//	fmt.Println("after goexit") // 这句不会执行
	//}()
	fmt.Println("编译器名", runtime.Compiler)
	fmt.Println("目标操作系统", runtime.GOOS)
	fmt.Println("目标系统架构", runtime.GOARCH)
	fmt.Println("当前系统使用的 cpu 核数", runtime.NumCPU())
	fmt.Println("goroot的环境变量", runtime.GOROOT())
	fmt.Println("go的版本", runtime.Version())
	//fmt.Println("go的版本", runtime.GC())
	fmt.Println("正在执行和排队的任务总数", runtime.NumGoroutine())
	fmt.Println("内存采样频率", runtime.MemProfileRate) // 默认 512kb 的内存分配一个样本，一般不建议配置
	fmt.Println("当前进程调用c方法的次数", runtime.NumCgoCall())
	runtime.SetCPUProfileRate(1) // 设置 cpu profile 记录的速率(平均 hz次/秒)，hz<=0表示关闭 profile 记录，大多数

	var h *human = &human{name: "q"}
	runtime.SetFinalizer(h, func(h *human) { // 绑定变量，在垃圾回收时进行监听
		println("垃圾回收了")
	})
	fmt.Println(h)
	runtime.GC() // 立即执行一次垃圾回收，这个需要了解垃圾回收的机制,对于不再使用的变量才会回收

	// 内存申请和分配统计
	var _ *human = &human{name: "q"}
	memStatus := runtime.MemStats{}
	runtime.ReadMemStats(&memStatus)
	fmt.Printf("申请的内存:%d\n", memStatus.Mallocs)
	fmt.Printf("释放的内存次数:%d\n", memStatus.Frees)
	fmt.Printf("指针查找的次数:%d\n", memStatus.Lookups)

	//runtime.Breakpoint()

	memProfileRecord := runtime.MemProfileRecord{}
	fmt.Printf("正在使用的字节数:%d\n", memProfileRecord.InUseBytes())
	fmt.Printf("正在使用的对象数:%d\n", memProfileRecord.InUseObjects())
	fmt.Printf("分配字节数:%d\n", memProfileRecord.AllocBytes)
	fmt.Printf("空闲的字节数:%d\n", memProfileRecord.FreeBytes)

	buf := make([]byte, 10000)
	go func() {
		println(11)
	}()
	go func() {
	}()
	runtime.Stack(buf, true) // 程序调用go协程的栈踪迹历史
	fmt.Println(string(buf))
	go test12()
	test10()
	time.Sleep(time.Second)
	test11()
}

func test10() {
	fmt.Println("==============")
	pc, file, line, ok := runtime.Caller(1)
	fmt.Printf("调用处的指针(标识号): %v\n", pc)
	fmt.Printf("调用处的文件: %s\n", file)
	fmt.Printf("调用处的行号: %d\n", line)
	fmt.Printf("无法返回信息返回false: %t\n", ok)
}
func test11() {
	fmt.Println("==============")
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs) // 获取与当前堆栈记录相关链的调用栈踪迹
	for _, pc := range pcs[:i] {
		funcPc := runtime.FuncForPC(pc)   // 获取标识调用栈标识符pc对应的调用栈
		name := funcPc.Name()             // 获取调用栈所调用的函数的名字
		file, line := funcPc.FileLine(pc) // 获取调用栈所调用的函数的所在的源文件名和行号
		fmt.Println(name, file, line)
		//funcPc.Entry() // 获取该调用栈的调用栈标识符
	}
}

func test12() {
	fmt.Println("==============")
	// 获取活跃的go协程的堆栈profile以及记录个数
	var st []runtime.StackRecord
	n, ok := runtime.GoroutineProfile(st)
	fmt.Println(st)
	fmt.Println(n, ok)
}
