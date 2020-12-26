package mycode

// https://leetcode.com/problems/decode-ways/discuss/30358/Java-clean-DP-solution-with-explanation
// dp[i]: num of ways to decode substring of first i chars
// dp[i] = f1 * dp[i-1] + f2 * dp[i-2]
// f1 = true iff '1' <= s[i-1] <= '9'
// f2 = true iff '10' <= s[i-2:i] <= '26'
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		f1, f2 := 0, 0
		if s[i-1] >= '1' && s[i-1] <= '9' {
			f1 = 1
		}
		if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			f2 = 1
		}
		dp[i] = f1*dp[i-1] + f2*dp[i-2]
	}

	return dp[n]
}
