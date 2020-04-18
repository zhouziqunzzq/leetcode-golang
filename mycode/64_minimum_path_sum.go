package mycode

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	for i, g := range grid {
		for j, _ := range g {
			if i == 0 && j == 0 {
				// base case 1
				continue
			} else if i == 0 {
				// base case 2
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				// base case 3
				grid[i][j] += grid[i-1][j]
			} else {
				// general cases
				if grid[i-1][j] < grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
