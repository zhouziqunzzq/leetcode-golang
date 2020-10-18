package mycode

// dp[i, k] = max(dp[i-1, k], prices[i] - prices[j] + dp[j-1, k-1])
// i: considering day 1...i
// k: with k transactions
// compact form
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// if 2k > n, we can trade everyday
	if 2*k > len(prices) {
		ans := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				ans += prices[i] - prices[i-1]
			}
		}
		return ans
	}

	dp := make([]int, k+1)
	min := make([]int, k+1)
	for i := range min {
		min[i] = prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for ck := 1; ck <= k; ck++ {
			min[ck] = minInt(min[ck], prices[i]-dp[ck-1])
			dp[ck] = maxInt(dp[ck], prices[i]-min[ck])
		}
	}

	return dp[k]
}
