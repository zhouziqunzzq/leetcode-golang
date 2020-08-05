package mycode

import "math"

// https://leetcode.com/problems/string-compression-ii/discuss/757363/Java-Easy-and-fast-DP-solution-with-comments-50ms-Space-O(n2)
func getLengthOfOptimalCompression(s string, k int) int {
	// opt[i][k] = the min length for s[:i] after at most k deletions
	var opt [110][110]int
	n := len(s)

	// init table
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			opt[i][j] = math.MaxInt32
		}
	}

	// base case
	opt[0][0] = 0

	// fill the table
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			same, diff := 0, 0
			// Option 1: keep s[i]
			for l := i; l >= 1; l-- {
				// try to remove all chars that's different from s[i]
				if s[l-1] == s[i-1] {
					same++
				} else {
					diff++
				}
				if j-diff >= 0 {
					opt[i][j] = minInt(
						opt[i][j],
						opt[l-1][j-diff]+1+getLengthOfOptimalCompressionHelper(same),
					)
				}
			}
			// Option 2: del s[i]
			if j > 0 {
				opt[i][j] = minInt(opt[i][j], opt[i-1][j-1])
			}
		}
	}

	return opt[n][k]
}

func getLengthOfOptimalCompressionHelper(cnt int) int {
	if cnt >= 100 {
		return 3
	}
	if cnt >= 10 {
		return 2
	}
	if cnt >= 2 {
		return 1
	}
	return 0
}
