package mycode

import "math"

func maxProfit1(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, p := range prices {
		minPrice = minInt(minPrice, p)
		maxProfit = maxInt(maxProfit, p-minPrice)
	}

	return maxProfit
}
