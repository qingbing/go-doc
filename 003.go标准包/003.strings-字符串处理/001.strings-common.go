package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"
)

func isAscii(r rune) bool {
	if r > 100 {
		return true
	}
	return false
}

func main() {
	s10 := "hello, love LVE love 中国"
	s11 := "llo"
	r10 := rune('中') // 20013
	b10 := byte('e')
	prefix := "hell"
	suffix := "中国"
	sLove := "love"

	fmt.Println(r10)
	// 长度
	fmt.Printf("(%s) len %d\n", s10, len(s10))
	// 包含关系
	fmt.Printf("(%s) Contains (%s)： %t\n", s10, s11, strings.Contains(s10, s11))
	fmt.Printf("(%s) ContainsAny (%s)： %t\n", s10, s11, strings.ContainsAny(s10, s11))
	fmt.Printf("(%s) ContainsRune (%c)： %t\n", s10, r10, strings.ContainsRune(s10, r10))
	fmt.Printf("(%s) HasPrefix (%s)： %t\n", s10, prefix, strings.HasPrefix(s10, prefix))
	fmt.Printf("(%s) HasSuffix (%s)： %t\n", s10, suffix, strings.HasSuffix(s10, suffix))
	fmt.Printf("(%s) Count (%s)： %d\n", s10, sLove, strings.Count(s10, sLove))
	// 查找
	fmt.Printf("(%s) Index (%s)： %d\n", s10, sLove, strings.Index(s10, sLove))
	fmt.Printf("(%s) IndexAny (%s)： %d\n", s10, sLove, strings.IndexAny(s10, sLove))
	fmt.Printf("(%s) IndexByte (%b)： %d\n", s10, b10, strings.IndexByte(s10, b10))
	fmt.Printf("(%s) IndexRune (%c)： %d\n", s10, r10, strings.IndexRune(s10, r10))
	fmt.Printf("(%s) IndexFunc (%v)： %d\n", sLove, isAscii, strings.IndexFunc(sLove, isAscii))
	// 转化
	fmt.Printf("(%s) ToLower： %s\n", s10, strings.ToLower(s10))
	fmt.Printf("(%s) ToUpper： %s\n", s10, strings.ToUpper(s10))
	fmt.Printf("(%s) ToTitle： %s\n", s10, strings.ToTitle(s10))
	fmt.Printf("(%s) ToLowerSpecial： %s\n", s10, strings.ToLowerSpecial(unicode.SpecialCase{}, s10))
	fmt.Printf("(%s) ToUpperSpecial： %s\n", s10, strings.ToUpperSpecial(unicode.SpecialCase{}, s10))
	fmt.Printf("(%s) ToTitleSpecial： %s\n", s10, strings.ToTitleSpecial(unicode.SpecialCase{}, s10))
	fmt.Printf("(%s) ToValidUTF8： %s\n", s10, strings.ToValidUTF8(s10, "llo"))

	// 对比
	s20 := "aa"
	s21 := "bb"
	s22 := "Bb"
	fmt.Printf("(%s) Compare (%s): %d\n", s20, s21, strings.Compare(s20, s21))
	fmt.Printf("(%s) Compare (%s): %d\n", s21, s20, strings.Compare(s21, s20))
	fmt.Printf("(%s) Compare (%s): %d\n", s21, s22, strings.Compare(s21, s22))
	fmt.Printf("(%s) EqualFold (%s): %t\n", s21, s22, strings.EqualFold(s21, s22))

	// trim
	s30 := "   aahello bb ccaa   "
	fmt.Printf("(%s) Trim： %s\n", s30, strings.Trim(s30, "ah"))
	fmt.Printf("(%s) TrimLeft： %s\n", s30, strings.TrimLeft(s30, "a"))
	fmt.Printf("(%s) TrimRight： %s\n", s30, strings.TrimRight(s30, "ac"))
	fmt.Printf("(%s) TrimPrefix： %s\n", s30, strings.TrimPrefix(s30, "aah"))
	fmt.Printf("(%s) TrimSuffix： %s\n", s30, strings.TrimSuffix(s30, "a"))
	fmt.Printf("(%s) Trim： %s\n", s30, strings.Trim(s30, " "))
	fmt.Printf("(%s) TrimSpace： %s\n", s30, strings.TrimSpace(s30))
	// 分割
	s40 := "this|is|a|testing"
	fmt.Println(strings.Split(s40, "|"))
	fmt.Println(strings.SplitN(s40, "|", 3))
	fmt.Println(strings.SplitAfter(s40, "|"))
	fmt.Println(strings.SplitAfterN(s40, "|", 3))
	fmt.Println(strings.Cut("120.0.0.1:8080", ":")) // 按照 sep 分割成两段
	// 字符串组合
	fmt.Println(strings.Join([]string{"i", "am", "qingbing"}, " "))
	// 字符串重复
	fmt.Println(strings.Repeat("hello", 3))

	// 替换
	s50 := "abcdabcd"
	fmt.Println(strings.ReplaceAll(s50, "abc", "ok"))
	fmt.Println(strings.Replace(s50, "abc", "ok", 1))
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'a' { // 用 @ 替换 a
			return '@'
		}
		if r == 'b' { // 删除b
			return 0
		}
		return r
	}, s50))
	replacer := strings.NewReplacer("<", "(", ">", ")")
	s51 := "i like <book>><"
	fmt.Println(replacer.Replace(s51))

	// 定义高效的字符串阅读器
	sReader := strings.NewReader("this is a string")
	fmt.Println(sReader.Len())
	i := 0
	for {
		b, err := sReader.ReadByte()
		if err == nil {
			fmt.Printf("第 %d 个字符: %c\n", i, b)
			i++
		} else if err == io.EOF {
			fmt.Println("读取完毕")
			break
		} else {
			log.Fatal("字符串读取失败")
		}
	}
}
