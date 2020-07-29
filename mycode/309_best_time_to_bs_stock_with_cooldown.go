package mycode

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking)
func maxProfit309(prices []int) int {
	sold := 0
	hold := math.MinInt32
	rest := 0

	for _, p := range prices {
		prevSold := sold
		sold = hold + p
		hold = maxInt(hold, rest-p)
		rest = maxInt(rest, prevSold)
	}

	return maxInt(rest, sold)
}
