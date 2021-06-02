package mycode

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	ans := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 && !visited[i][j] {
				ans = maxInt(ans, maxAreaOfIslandDFS(grid, visited, i, j, d))
			}
		}
	}
	return ans
}

func maxAreaOfIslandDFS(grid [][]int, visited [][]bool, x, y int, d [][]int) int {
	if grid[x][y] == 0 || visited[x][y] {
		return 0
	}

	m, n := len(grid), len(grid[0])
	cnt := 1
	visited[x][y] = true

	for _, dd := range d {
		newX, newY := x+dd[0], y+dd[1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n &&
			grid[newX][newY] == 1 && !visited[newX][newY] {
			cnt += maxAreaOfIslandDFS(grid, visited, newX, newY, d)
		}
	}

	return cnt
}
