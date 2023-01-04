

pprof使用
编写一个简单的应用程序，使用pprof.StartCPUProfile和pprof.StopCPUProfile对CPU信息进行采样。

package mainimport (   "flag"   "log"   "os"   "runtime/pprof"   "fmt")// 斐波纳契数列func Fibonacci() func() int {   back1, back2 := 1, 1   return func() int {      //重新赋值      back1, back2 = back2, (back1 + back2)      return back1   }}func count(){    a := 0;   for i := 0; i < 10000000000; i++ {      a = a + i   }}var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")func main() {   flag.Parse()   if *cpuprofile != "" {      f, err := os.Create(*cpuprofile)      if err != nil {         log.Fatal(err)      }      pprof.StartCPUProfile(f)      defer f.Close()   }   fibonacci := Fibonacci()   for i := 0; i < 100; i++ {      fmt.Println(fibonacci())   }   count()   defer pprof.StopCPUProfile()}
进行运行时信息采样时，可以指定不同的采样参数：
--cpuprofile：指定CPU概要文件的保存路径
--blockprofile：指定程序阻塞概要文件的保存路径。
--blockprofilerate：定义其值为n，指定每发生n次Goroutine阻塞事件时，进行一次取样操作。
--memprofile：指定内存概要文件的保存路径。
--memprofilerate：定义其值为n，指定每分配n个字节的堆内存时，进行一次取样操作

