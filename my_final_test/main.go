package main

import (
	"fmt"
	"strconv"
)

// 可以由调用方决定类型的回调函数
type SumCallBack[T any] func([]T) T

func main() {
	result := newArray(func(array []int) int {
		sum := 0
		for _, value := range array {
			sum += value
		}
		return sum
	})
	fmt.Printf("result:%v\n", result)
}

func newArray(myCallback SumCallBack[int]) string {
	array := []int{1, 2, 3, 4, 6}
	prepare := make([]int, len(array))
	for _, value := range array {
		prepare = append(prepare, value)
	}
	result := myCallback(prepare)
	fmt.Printf("回调函数的结果是:%v\n", result)
	divide := result / 2
	return strconv.Itoa(divide)
}

func newMap(myCallback SumCallBack[int]) {
	myMap := make(map[string]struct{})
	myMap["zhubo"] = struct{}{}

}
