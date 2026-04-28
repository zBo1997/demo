package main

import "fmt"

// Add
//
//	@param a 指针
//	@param args
//	@return int
//
// 这个方法接收一个整数指针和可变数量的整数参数，将这些整数累加到指针指向的值上，并返回累加后的结果
func Add(a *int, args ...int) int {
	result := a
	for _, v := range args {
		*result += v
	}
	return *result
}

// 不带使用指针的版本
func AddNormal(a int, args ...int) int {
	result := a
	for _, v := range args {
		result += v
	}
	return result
}

func main() {
	//做一个测试，看看那个版本的性能最好？
	var count = 1000
	var p = 1
	for i := 0; i < count; i++ {
		fmt.Println(Add(&p, 1))
	}
	// for i := 0; i < count; i++ {
	// 	fmt.Println(AddNormal(p, i))
	// }

}
