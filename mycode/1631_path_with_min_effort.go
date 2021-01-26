package mycode

// https://leetcode.com/problems/path-with-minimum-effort/discuss/909002/JavaPython-3-3-codes%3A-Binary-Search-Bellman-Ford-and-Dijkstra-w-brief-explanation-and-analysis.
// Solution 1: Bin-search + BFS/DFS <- implemented
// Solution 2: SPFA (TLE)
// Solution 3: Dijkstra
func minimumEffortPath(heights [][]int) int {
	l, r := 0, int(1e6)
	m, n := len(heights), len(heights[0])
	d := []int{0, 1, 0, -1, 0}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for l < r {
		m := l + (r-l)/2
		if minimumEffortPathHelper(heights, visited, d, m, 0, 0) {
			// m is large enough, we can try smaller values
			r = m
		} else {
			// m is too small, try larger values
			l = m + 1
		}

		// reset visited
		for i, vis := range visited {
			for j := range vis {
				visited[i][j] = false
			}
		}
	}

	return l
}

func minimumEffortPathHelper(h [][]int, visited [][]bool, d []int, k, x, y int) bool {
	m, n := len(h), len(h[0])

	if x == m-1 && y == n-1 {
		return true
	}

	visited[x][y] = true
	for i := 0; i < 4; i++ {
		nx, ny := x+d[i], y+d[i+1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n &&
			!visited[nx][ny] &&
			absInt(h[nx][ny]-h[x][y]) <= k {
			if minimumEffortPathHelper(h, visited, d, k, nx, ny) {
				return true
			}
		}
	}
	return false
}
