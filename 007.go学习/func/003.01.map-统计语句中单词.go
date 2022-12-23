package main

import (
	"fmt"
	"strings"
)

/*
map-统计语句中单词
*/
func countWords(sentence string) (out map[string]int){
	out = make(map[string]int)
	words := strings.Fields(sentence);
	for _, word := range words {
		if _, has:=out[word]; has {
			out[word]++
		}else{
			out[word] = 1
		}
	}
	fmt.Println(words)
	return
}

func main(){
	sentence := "This is a app , I like app !"
	counts := countWords(sentence)
	fmt.Println("sentence ===> ", sentence)
	fmt.Println("Counts ===> ", counts)
}

