package main

import (
	"fmt"
)

func main() {
	array := []int{2, 10, 3, 1, 5, 7}
	quickSort(array)
	fmt.Println(array)
}

func partition(array []int) int {
	priv := array[len(array)-1]
	i := 0
	for j := 0; j < len(array)-1; j++ {
		if array[j] < priv {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[len(array)-1] = array[len(array)-1], array[i]
	return i
}

func quickSort(array []int) {
	if len(array) == 0 {
		return
	}
	privIndex := partition(array)

	quickSort(array[:privIndex])
	quickSort(array[privIndex+1:])
}
