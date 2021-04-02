package mycode

// reduce to 0-1 Knapsack problem.
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j][k]: max size of the subset we can get considering only strs[:i] with at most j '0's and k '1's
	// here we can omit i since dp[i] depends only on dp[i-1]
	var dp [101][101]int

	for _, s := range strs {
		// count 0s and 1s
		cnt0, cnt1 := 0, 0
		for _, c := range s {
			if c == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}

		// update dp
		// NOTE: we do this with descending j and k to prevent from messing up dp[i-1]
		for j := m; j >= cnt0; j-- {
			for k := n; k >= cnt1; k-- {
				// choice 1: don't take the ith str so dp[i][j][k] = dp[i-1][j][k]
				// choice 2: take the ith str and dp[i][j][k] = dp[i-1][j-cnt0][k-cnt1]+1
				dp[j][k] = maxInt(dp[j][k], dp[j-cnt0][k-cnt1]+1)
			}
		}
	}

	return dp[m][n]
}
