package mycode

func uniquePathsIII(grid [][]int) int {
	rst := 0
	nonObstacles := 0
	startX, startY := 0, 0
	for x, g := range grid {
		for y, f := range g {
			if f >= 0 {
				nonObstacles++
			}
			if f == 1 {
				startX, startY = x, y
			}
		}
	}

	uniquePathsIIIHelper(grid, startX, startY, nonObstacles, &rst)

	return rst
}

func uniquePathsIIIHelper(grid [][]int, x, y, remain int, cnt *int) {
	if grid[x][y] < 0 || remain == 0 {
		return
	}

	if grid[x][y] == 2 && remain == 1 {
		*cnt++
		return
	}

	origin := grid[x][y]
	grid[x][y] = -4 // mark as visited

	xDiff := []int{0, 0, 1, -1}
	yDiff := []int{1, -1, 0, 0}
	for i := 0; i < 4; i++ {
		xNext := x + xDiff[i]
		yNext := y + yDiff[i]
		if xNext < 0 || xNext >= len(grid) || yNext < 0 || yNext >= len(grid[0]) {
			continue
		}
		uniquePathsIIIHelper(grid, xNext, yNext, remain-1, cnt)
	}

	grid[x][y] = origin
}
