package main

import "fmt"

func main() {
	var str1 = "hello"
	//新建立一个hello对象 使其内存地址不一样
	var str2 = "hello"

	fmt.Printf("str1: %p, str2: %p\n", &str1, &str2)

	if str1 == str2 {
		fmt.Println("str1 and str2 are equal")
	} else {
		fmt.Println("str1 and str2 are not equal")
	}
}
