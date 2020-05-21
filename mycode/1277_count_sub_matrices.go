package mycode

func countSquares(matrix [][]int) int {
	// dp[i][j] is the edge length of the largest square with right bottom corner at (i, j)
	// we can further improve the space cost by using matrix as dp[][]
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	ans := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = matrix[i][j]
			if i > 0 && j > 0 && dp[i][j] > 0 {
				dp[i][j] = minInt3(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			}
			ans += dp[i][j]
		}
	}

	return ans
}
