package mycode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/135704/Detail-explanation-of-DP-solution
// dp[i, k] = max(dp[i-1, k], prices[i] - prices[j] + dp[j-1, k-1])
// i: considering day 1...i
// k: with k transactions
// compact form
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, 3)
	min := make([]int, 3)
	min[0], min[1], min[2] = prices[0], prices[0], prices[0]

	for i := 1; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			min[k] = minInt(min[k], prices[i]-dp[k-1])
			dp[k] = maxInt(dp[k], prices[i]-min[k])
		}
	}

	return dp[2]
}
