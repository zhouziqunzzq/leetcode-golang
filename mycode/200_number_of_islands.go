package mycode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	ans := 0
	for i, g := range grid {
		for j, b := range g {
			if b == '1' {
				ans++
				islandDFS(grid, i, j)
			}
		}
	}

	return ans
}

func islandDFS(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	// up
	if x-1 >= 0 && grid[x-1][y] == '1' {
		islandDFS(grid, x-1, y)
	}
	// down
	if x+1 < len(grid) && grid[x+1][y] == '1' {
		islandDFS(grid, x+1, y)
	}
	// left
	if y-1 >= 0 && grid[x][y-1] == '1' {
		islandDFS(grid, x, y-1)
	}
	// right
	if y+1 < len(grid[0]) && grid[x][y+1] == '1' {
		islandDFS(grid, x, y+1)
	}
}
