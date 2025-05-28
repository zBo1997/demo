package main

//轮转数组
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(result, newNums)
}

func main() {
	k := 3
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, k)
}
