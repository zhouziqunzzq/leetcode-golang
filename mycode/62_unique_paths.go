package mycode

// calculate combination of C(m-1+n-1, m-1) with dp
func uniquePaths(m int, n int) int {
	opt := make([]int, n)

	// base case
	for i, _ := range opt {
		opt[i] = 1
	}

	// fill the table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			opt[j] += opt[j-1]
		}
	}

	return opt[n-1]
}
