package mycode

import "strconv"

// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/discuss/961446/Detailed-Thought-Process-with-Steps-Example-or-Java-8-1-Liner
func concatenatedBinary(n int) int {
	mod := 1_000_000_007
	sum := 0
	for i := 1; i <= n; i++ {
		sum = ((sum << len(strconv.FormatInt(int64(i), 2))) + i) % mod
	}

	return sum
}
