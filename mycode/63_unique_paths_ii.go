package mycode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// fill 1st row
	t := 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			t = 0
		}
		dp[0][j] = t
	}

	// fill 1st col
	t = 1
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			t = 0
		}
		dp[i][0] = t
	}

	// fill the table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) * (1 - obstacleGrid[i][j])
		}
	}

	return dp[m-1][n-1]
}
