package mycode

// https://leetcode.com/problems/rotate-image/solution/
// Solution 1: Rotate groups of four cells
// Solution 2: Transpose + Reflect
func rotate(matrix [][]int) {
	transpose48(matrix)
	reflect48(matrix)
}

func transpose48(m [][]int) {
	for i := range m {
		for j := i; j < len(m); j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

func reflect48(m [][]int) {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			m[i][j], m[i][n-j-1] = m[i][n-j-1], m[i][j]
		}
	}
}
