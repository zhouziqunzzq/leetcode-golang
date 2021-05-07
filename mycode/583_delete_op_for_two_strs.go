package mycode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, row := range dp {
		for j := range row {
			if i == 0 || j == 0 { // base case
				dp[i][j] = i + j
			} else {
				if word1[i-1] == word2[j-1] { // no need to delete
					dp[i][j] = dp[i-1][j-1]
				} else { // must delete
					dp[i][j] = 1 + minInt(dp[i][j-1], dp[i-1][j])
				}
			}
		}
	}

	return dp[m][n]
}
