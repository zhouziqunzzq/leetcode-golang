package mycode

// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/discuss/288520/Longest-Path-in-DAG
// Use BFS to do topological sort and get the longest path in DAG.
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ds := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	inDegree := make([][]int, m)
	for j := range inDegree {
		inDegree[j] = make([]int, n)
	}
	for i, r := range inDegree {
		for j := range r {
			for _, d := range ds {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if matrix[ni][nj] > matrix[i][j] {
						inDegree[ni][nj]++
					}
				}
			}
		}
	}

	// initialize Queue with nodes that have zero in-degree
	q := make([][]int, 0)
	for i, r := range inDegree {
		for j, ind := range r {
			if ind == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	ans := 0
	for len(q) > 0 {
		qn := len(q)
		// expand queue level by level
		for k := 0; k < qn; k++ {
			i, j := q[k][0], q[k][1]
			for _, d := range ds {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if matrix[ni][nj] > matrix[i][j] {
						// del edge (i, j) -> (ni, nj)
						inDegree[ni][nj]--
						if inDegree[ni][nj] == 0 {
							q = append(q, []int{ni, nj})
						}
					}
				}
			}
		}
		q = q[qn:]
		ans++
	}

	return ans
}
