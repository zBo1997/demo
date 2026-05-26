package main

import (
	"fmt"
	"sort"
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

	resultMap := newMap(func(array []string) string {
		result := ""
		//把map按照key排序
		sortedArray := make([]string, len(array))
		copy(sortedArray, array)
		sort.Strings(sortedArray)
		for _, value := range sortedArray {
			result += value + " "
		}
		return result
	})
	fmt.Printf("resultMap:%v\n", resultMap)
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

func newMap(myCallback SumCallBack[string]) string {
	myMap := make(map[string]struct{})
	myMap["zhubo5"] = struct{}{}
	myMap["zhubo1"] = struct{}{}
	myMap["zhubo3"] = struct{}{}
	myMap["zhubo2"] = struct{}{}
	prepare := make([]string, 0, len(myMap))
	for key := range myMap {
		prepare = append(prepare, key)
	}
	result := myCallback(prepare)
	return result
}
