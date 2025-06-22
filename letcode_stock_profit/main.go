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

	var maxProfit = 0
	var currentProfit = 0
	for i := 0; i < len(prices); i++ {
		currentPrice := prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > currentPrice {
				currentProfit = prices[j] - currentPrice
				maxProfit = max(maxProfit, currentProfit)
			}
		}
	}
	return maxProfit
}

func max(a,b int) int {
    if a > b {
        return a
    }

    return b
}


func main() {
	nums := []int{7,1,5,3,6,4}
	result := maxProfit(nums)
	fmt.Println(result)
}

