package main

import (
	"fmt"
)


//股票最大利润
func maxProfit(prices []int) int {
	//确定边界条件
	if len(prices) < 2 {
        return 0
    }

	//确定最小值
	minPrice := 0
	minIndex := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice || minPrice == 0 {
			minPrice = prices[i]
			minIndex = i
		}
	}

	//如果最小值为数组最后一个元素，说明没有买入时机
	if minPrice == prices[len(prices) - 1] {
		return 0
	}

	fmt.Println(minPrice)

	maxProfit := 0
	for i := minIndex ; i < len(prices); i++ {
		if minPrice < prices[i] {
			currentProfit := prices[i] - minPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	return maxProfit
}

func main() {
	nums := []int{7,1,5,3,6,4}
	result := maxProfit(nums)
	fmt.Println(result)
}

