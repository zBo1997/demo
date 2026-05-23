package main

import "fmt"

// 最长有效括号子串为支持 英文括号 中括号 花括号
func main() {
	s := "({})[]{}"
	test := s[0:1]
	fmt.Println(test)
	result := longestValidParentheses(s)
	println("最长有效括号子串为:", result)
}

// 最长有效括号子串为支持 英文括号 中括号 花括号
func longestValidParentheses(s string) string {
	if len(s) == 0 {
		return ""
	}
	startIndex := 0
	maxLen := 0
	n := len(s)
	stack := []int{-1}

	for i := range n {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i) // 重新设置基准索引
			} else {
				currLen := i - stack[len(stack)-1]
				if currLen > maxLen {
					maxLen = currLen
					startIndex = stack[len(stack)-1] + 1
				}
			}
		}
	}
	if maxLen == 0 {
		return ""
	}
	return s[startIndex : startIndex+maxLen]
}
