package mycode

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	x, y := 0, 0
	// 0: right, 1: down, 2: left, 3: up
	dir := 0
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for i := 1; i <= n*n; i++ {
		m[x][y] = i
		nx, ny := x+dx[dir], y+dy[dir]
		if (nx < 0 || nx >= n || ny < 0 || ny >= n) || (m[nx][ny] != 0) {
			dir = (dir + 1) % 4
			nx, ny = x+dx[dir], y+dy[dir]
		}
		x, y = nx, ny
	}

	return m
}
