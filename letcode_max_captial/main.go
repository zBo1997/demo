package main

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。

func maxArea(height []int) int {
	anws := 0
	left := 0
	right := len(height) - 1
	// 双指针法 指针没有相遇的时候
	for left < right {
		area := (right - left) * min(height[left], height[right])
		anws = max(anws, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return anws
}

func main() {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(s)
	println(result) // Output: 5
}
