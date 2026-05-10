package main

import "reflect"

// 我们有一组数据 含有字母 正整数字 浮点数 只对数字快速排序
// 双指针 是一个什么递归
func quick_sort(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	//pivot是指针 选取数组中间的元素作为基准元素的值
	pivot := arr[len(arr)/2]
	left := []float64{}
	right := []float64{}
	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	// 递归调用 quick_sort 对 left 和 right 进行排序，并将结果与 pivot 连接起来
	return append(append(quick_sort(left), pivot), quick_sort(right)...)
}

// 浮点数
func extract_number(data []interface{}) []float64 {
	var nums []float64
	for _, item := range data {
		//查看类型
		val := reflect.ValueOf(item)
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			nums = append(nums, float64(val.Int()))
		case reflect.Float32, reflect.Float64:
			nums = append(nums, val.Float())
		}
	}
	return nums
}

func main() {
	data := []interface{}{"a", 5, 2.5, "b", 3, 4.5, "c"}
	nums := extract_number(data)
	sorted_nums := quick_sort(nums)
	println("排序后的数字：")
	for _, num := range sorted_nums {
		println(num)
	}
}
