package mycode

// https://leetcode.com/problems/cherry-pickup-ii/solution/
// dp[row][col1][col2]: max cherry two robots can pick starting at
// (row, col1) and (row, col2) respectively.
func cherryPickup(grid [][]int) int {
	var dp [71][71][71]int
	m, n := len(grid), len(grid[0]) // m rows, n cols

	for row := m - 1; row >= 0; row-- {
		for col1 := 0; col1 < n; col1++ {
			for col2 := 0; col2 < n; col2++ {
				// current cell
				rst := grid[row][col1]
				if col1 != col2 {
					rst += grid[row][col2]
				}

				// transition
				if row != m-1 {
					nextMax := 0
					for nextCol1 := col1 - 1; nextCol1 <= col1+1; nextCol1++ {
						for nextCol2 := col2 - 1; nextCol2 <= col2+1; nextCol2++ {
							if nextCol1 >= 0 && nextCol1 < n && nextCol2 >= 0 && nextCol2 < n {
								nextMax = maxInt(nextMax, dp[row+1][nextCol1][nextCol2])
							}
						}
					}
					rst += nextMax
				}

				dp[row][col1][col2] = rst
			}
		}
	}

	return dp[0][0][n-1]
}
