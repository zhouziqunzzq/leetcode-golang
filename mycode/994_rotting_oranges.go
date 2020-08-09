package mycode

func orangesRotting(grid [][]int) int {
	q := make([]struct {
		x int
		y int
	}, 0)
	freshCnt := 0

	for i, g := range grid {
		for j, o := range g {
			if o == 2 {
				q = append(q, struct {
					x int
					y int
				}{x: i, y: j})
			} else if o == 1 {
				freshCnt++
			}
		}
	}

	days := 0
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for len(q) > 0 {
		if freshCnt == 0 {
			return days
		}
		days++
		curSize := len(q)
		for i := 0; i < curSize; i++ {
			o := q[0]
			q = q[1:]
			for _, d := range dirs {
				x := o.x + d[0]
				y := o.y + d[1]
				if x < 0 || y < 0 ||
					x >= len(grid) || y >= len(grid[0]) ||
					grid[x][y] == 0 || grid[x][y] == 2 {
					continue
				}
				grid[x][y] = 2
				q = append(q, struct {
					x int
					y int
				}{x: x, y: y})
				freshCnt--
			}
		}
	}

	if freshCnt == 0 {
		return 0
	} else {
		return -1
	}
}
