
- 格式化输出
  - %q: 以 go 语言格式显示字符从，默认带有 "" 符号
  - %v: 显示对应数据的详细信息
- 变量存储
  - 等号左边的变量， 代表变量所指向的内存空间 (写)
  - 等号右边的变量， 代表变量内存空间存储的数据值 (读)
- 指针的函数传参
  - 传地址(传引用): 将实参的地址值作为函数的参数，最大的作用时，可以在被调用函数修改调用处的值
  - 传值(传数据): 将实参的值拷贝一份给形参


- make 只能用于创建 slice、map、channel, 并且返回的是一个有初始值（非零）的对象
- 函数参数传引用: slice, map, pointer
- len 函数可针对 Array、Pointer to array、Slice、Map、String、Channel
- cap 函数可针对 Array、Pointer to array、Slice、Channel

