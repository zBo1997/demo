package main

import (
	"fmt"
	"strconv"
)

type myCallback func(param ...interface{}) string

func main() {
	result := newArray(func(param ...interface{}) string {
		var result = 0
		for _, value := range param {
			if num, ok := value.(int); ok {
				result += num
			}
		}
		return strconv.Itoa(result)
	})
	fmt.Printf("result:%v\n", result)
}

func newArray(myCallback myCallback) string {
	array := []int{1, 2, 3, 4, 5}
	prepare := make([]interface{}, len(array))
	for _, value := range array {
		prepare = append(prepare, value)
	}
	return myCallback(prepare...)
}

func newMap(myCallback myCallback) {
	myMap := make(map[string]struct{})
	myMap["zhubo"] = struct{}{}

}
