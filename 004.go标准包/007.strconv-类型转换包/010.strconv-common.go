package main

import (
	"fmt"
	"strconv"
)

func main() {
	i1, _ := strconv.Atoi("11")
	fmt.Printf("字符串转换成数字: %d\n", i1)
	fmt.Printf("数字转换成字符串: %s\n", strconv.Itoa(111))

	b2, _ := strconv.ParseBool("true")
	f2, _ := strconv.ParseFloat("3.1415", 64)
	i2, _ := strconv.ParseInt("-42", 10, 64)
	u2, _ := strconv.ParseUint("42", 10, 64)
	fmt.Printf("true: %t\n", b2)
	fmt.Printf("3.1415: %f\n", f2)
	fmt.Printf("-42: %d\n", i2)
	fmt.Printf("42: %d\n", u2)
	// error
	u3, err := strconv.ParseInt("17", 32, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		println("转换成功", u3)
	}

	// 转换成字符串
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(61, 8)) // 61 表示成 8进制的字符串
	// float到string: fmt 参考字符串占位符; prec 为精准控制位(对于'e'、'E'、'f'、'x'和'X'，它是小数点后的位数; 对于 'g' 和 'G' 它是有效数字的最大数量,尾随零被删除; -1 为特殊精度 使用最少的位数); bitSize 为32｜64
	f := strconv.FormatFloat(3.1415926, 'f', -1, 32)
	f32 := strconv.FormatFloat(3.1415926, 'e', -1, 32)
	f64 := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Printf("type:%T; 值:%[1]v\n", f)
	fmt.Printf("type:%T; 值:%[1]v\n", f32)
	fmt.Printf("type:%T; 值:%[1]v\n", f64)

	// Quote 系列
	var s4 string = "Hello, 世界"
	var r4 rune = 7
	var s42 string = "\"hello\""
	fmt.Printf("String: %s, Quote: %s\n", s4, strconv.Quote(s4))
	fmt.Printf("String: %s, QuoteToASCII: %s\n", s4, strconv.QuoteToASCII(s4))
	fmt.Printf("String: %s, QuoteToGraphic: %s\n", s4, strconv.QuoteToGraphic(s4))
	fmt.Printf("Rune: %d, QuoteRune: %s\n", r4, strconv.QuoteRune(r4))
	fmt.Printf("Rune: %d, QuoteRuneToASCII: %s\n", r4, strconv.QuoteRuneToASCII(r4))
	fmt.Printf("Rune: %d, QuoteRuneToGraphic: %s\n", r4, strconv.QuoteRuneToGraphic(r4))
	s43, _ := strconv.Unquote(s42)
	fmt.Printf("quote: %s, Unquote: %s\n", s42, s43)

	// QuotedPrefix
	s, err := strconv.QuotedPrefix("not a quoted string") // invalid syntax
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.QuotedPrefix("\"double-quoted string\" with trailing text") // "double-quoted string"
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.QuotedPrefix("`or backquoted` with more trailing text") // `or backquoted`
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.QuotedPrefix("'\u263a' is also okay") // '☺'
	fmt.Printf("%q, %v\n", s, err)

	// 判断
	fmt.Printf("IsPrint: %d, %t\n", '\u263a', strconv.IsPrint('\u263a'))
	fmt.Printf("IsGraphic: %d, %t\n", 97, strconv.IsGraphic(97))
	fmt.Printf("CanBackquote:  %t\n", strconv.CanBackquote("asf\nsdf")) // 有换行，为 false

	// 字符串追加
	s51 := []byte{97}
	fmt.Println(string(s51))
	s51 = strconv.AppendBool(s51, true)
	s51 = strconv.AppendInt(s51, 11, 8)
	s51 = strconv.AppendQuote(s51, "hello")
	fmt.Println(string(s51))
}
