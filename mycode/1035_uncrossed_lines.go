package mycode

// LCS
func maxUncrossedLines(A []int, B []int) int {
	opt := make([][]int, len(A)+1)
	for i := range opt {
		opt[i] = make([]int, len(B)+1)
	}

	for i := 1; i < len(opt); i++ {
		for j := 1; j < len(opt[i]); j++ {
			if A[i-1] == B[j-1] {
				opt[i][j] = opt[i-1][j-1] + 1
			} else {
				opt[i][j] = maxInt(opt[i][j-1], opt[i-1][j])
			}
		}
	}

	return opt[len(A)][len(B)]
}
