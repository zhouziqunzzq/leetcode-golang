package mycode

// https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/66140/My-concise-O(m%2Bn)-Java-solution
// O(m+n), search from top-right or bottom-left
// it's like searching in a BST
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1

	for r < m && c >= 0 {
		cur := matrix[r][c]
		if target > cur {
			r++
		} else if target < cur {
			c--
		} else {
			return true
		}
	}

	return false
}
