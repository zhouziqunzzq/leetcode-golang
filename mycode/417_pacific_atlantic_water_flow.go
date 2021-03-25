package mycode

var pacificAtlanticDPos = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return make([][]int, 0)
	}
	m, n := len(matrix), len(matrix[0])

	flags := make([][]int, m)
	for i := range flags {
		flags[i] = make([]int, n)
	}

	// dfs from pacific
	for y := 0; y < n; y++ {
		if flags[0][y]&1 == 0 {
			pacificAtlanticDFS(matrix, flags, 0, y, 'p')
		}
	}
	for x := 0; x < m; x++ {
		if flags[x][0]&1 == 0 {
			pacificAtlanticDFS(matrix, flags, x, 0, 'p')
		}
	}

	// dfs from atlantic
	for y := 0; y < n; y++ {
		if flags[m-1][y]&2 == 0 {
			pacificAtlanticDFS(matrix, flags, m-1, y, 'a')
		}
	}
	for x := 0; x < m; x++ {
		if flags[x][n-1]&2 == 0 {
			pacificAtlanticDFS(matrix, flags, x, n-1, 'a')
		}
	}

	ans := make([][]int, 0)

	for x, fs := range flags {
		for y, f := range fs {
			if f&1 > 0 && f&2 > 0 {
				t := make([]int, 2)
				t[0], t[1] = x, y
				ans = append(ans, t)
			}
		}
	}

	return ans
}

// flags: 1 - p visited, 2 - a visited
func pacificAtlanticDFS(matrix [][]int, flags [][]int, x, y int, from rune) {
	m, n := len(matrix), len(matrix[0])

	if from == 'p' {
		flags[x][y] |= 1
	} else {
		flags[x][y] |= 2
	}

	for _, d := range pacificAtlanticDPos {
		newX, newY := x+d[0], y+d[1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n && matrix[x][y] <= matrix[newX][newY] {
			if (from == 'p' && flags[newX][newY]&1 > 0) || (from == 'a' && flags[newX][newY]&2 > 0) {
				continue
			}
			pacificAtlanticDFS(matrix, flags, newX, newY, from)
		}
	}
}
