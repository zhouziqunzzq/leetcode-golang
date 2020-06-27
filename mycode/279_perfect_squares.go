package mycode

import "math"

// Solution 1: DP <--- implemented this
// Solution 2: math	TODO
// Solution 3: BFS	TODO
func numSquares(n int) int {
	opt := make([]int, n+1)

	// base case
	opt[0] = 0

	for i := 1; i <= n; i++ {
		opt[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			opt[i] = minInt(opt[i], opt[i-j*j]+1)
		}
	}

	return opt[n]
}
