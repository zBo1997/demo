package main

import (
	"strings"
)

// 最后一个单词的长度
func lengthOfLastWord(s string) int {
	// 按照空过进行分割多个单词数组
	parts := strings.Fields(s)
	// 返回数组最后一个单词的长度
	return len(parts[len(parts)-1])
}

func main() {
	s := "Hello World"
	result := lengthOfLastWord(s)
	println(result) // Output: 5
}
