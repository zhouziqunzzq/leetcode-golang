package mycode

type Pos1091 struct {
	x    int
	y    int
	dist int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	d := []int{-1, 0, 1}
	var visited [100][100]bool
	q := make([]Pos1091, 0)
	if grid[0][0] == 0 {
		q = append(q, Pos1091{0, 0, 1})
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur.x == n-1 && cur.y == n-1 {
			return cur.dist
		}

		if visited[cur.x][cur.y] {
			continue
		}

		visited[cur.x][cur.y] = true
		for i := range d {
			for j := range d {
				nextX, nextY := cur.x+d[i], cur.y+d[j]
				if nextX >= 0 && nextX < n &&
					nextY >= 0 && nextY < n &&
					!visited[nextX][nextY] &&
					grid[nextX][nextY] == 0 {
					q = append(q, Pos1091{nextX, nextY, cur.dist + 1})
				}
			}
		}
	}

	return -1
}
