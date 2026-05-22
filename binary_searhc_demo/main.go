package main

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 5
	result := binarySearch(arr, target)
	if result != -1 {
		println("元素", target, "在数组中的索引为:", result)
	} else {
		println("元素", target, "不在数组中")
	}
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
