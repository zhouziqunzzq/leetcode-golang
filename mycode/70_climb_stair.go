package mycode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		dp := make([]int, 2)
		dp[0], dp[1] = 1, 2
		for i := 2; i < n; i++ {
			dp = append(dp, dp[i-1]+dp[i-2])
		}
		return dp[n-1]
	}
}
