package mycode

// https://leetcode.com/problems/range-sum-query-2d-immutable/solution/
// Solution: in anology to 1-D case, compute a cumulative region sum
// w.r.t. the origin at 0,0.
type NumMatrix struct {
	dp [][]int // dp[i][j] = region sum from (0,0) to (i,j) inclusively
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for r, rArr := range matrix {
		for c, elem := range rArr {
			// principle of inclusion-exclusion
			dp[r+1][c+1] = dp[r+1][c] + dp[r][c+1] + elem - dp[r][c]
		}
	}

	return NumMatrix{dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// principle of inclusion-exclusion
	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
