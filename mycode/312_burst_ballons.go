package mycode

func maxCoins(nums []int) int {
	n := len(nums)
	b := make([]int, n+2)
	b[0], b[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		b[i] = nums[i-1]
	}

	// dp[i][j]: max coins one can get by bursting sub-array b[i:j]
	// dp[l][r] = max(b[l-1]*b[k]*b[r+1] + dp[l][k-1] + dp[k+1][r]) where k = l...r
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ {
		for l := 1; l <= n-length+1; l++ {
			r := l + length - 1
			for k := l; k <= r; k++ {
				cur := b[l-1]*b[k]*b[r+1] + dp[l][k-1] + dp[k+1][r]
				if cur > dp[l][r] {
					dp[l][r] = cur
				}
			}
		}
	}

	return dp[1][n]
}
