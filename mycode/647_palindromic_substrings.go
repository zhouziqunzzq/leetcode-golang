package mycode

func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		if s[0] == s[1] {
			return 3
		} else {
			return 2
		}
	}

	// dp[i][j]: s[i:j] is palindromic or not
	// dp[i][j] = true iff dp[i+1][j-1] == true && s[i] == s[j]
	// dp[i][j] = false o.w.
	// TODO: can be further optimized to take up only O(N) space
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	ans := 0

	// base case: len = 1
	for i := range dp {
		dp[i][i] = true
	}
	ans += n

	// base case: len = 2
	for i := 0; i < n-1; i++ {
		j := i + 1
		dp[i][j] = s[i] == s[j]
		if dp[i][j] {
			ans++
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i < n-(l-1); i++ {
			j := i + l - 1
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] {
				ans++
			}
		}
	}

	return ans
}
