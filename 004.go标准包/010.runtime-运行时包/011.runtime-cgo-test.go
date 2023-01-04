package main

/*
   #include <stdio.h>
   // 自定义一个c语言的方法
   static void myPrint(const char* msg) {
     printf("myPrint: %s", msg);
   }
*/
import "C" // 内部init算一次c调用

func main() {
	// 调用c方法
	C.myPrint(C.CString("Hello,C\n"))
	//println(runtime.NumCgoCall())
}
