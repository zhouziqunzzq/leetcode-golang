package mycode

import "strconv"
import "math"

// https://leetcode.com/problems/numbers-at-most-n-given-digit-set/solution/
// digit DP
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	k := len(s)
	// dp[i]: number of ways to write a valid (k-i)-digit number for n[i...k-1]
	// dp[i] = (num of d in D s.t. d < n[i]) * (|D| ** (k-i-1)) + I[n[i] in D?] * dp[i+1]
	dp := make([]int, k+1)
	dp[k] = 1

	// compute dp[k-1, ..., 0]
	for i := k - 1; i >= 0; i-- {
		for _, d := range digits {
			if d[0] < s[i] {
				dp[i] += int(math.Pow(float64(len(digits)), float64(k-i-1)))
			} else if d[0] == s[i] {
				dp[i] += dp[i+1]
			}
		}
	}

	// add number of ways to write a valid number with number of digits less than k
	for i := 1; i < k; i++ {
		dp[0] += int(math.Pow(float64(len(digits)), float64(i)))
	}

	return dp[0]
}
