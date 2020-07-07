package mycode

// Solution 1: DFS <-- implemented here
// Solution 2: math ref https://leetcode.com/problems/island-perimeter/discuss/95001/clear-and-easy-java-solution
func islandPerimeter(grid [][]int) int {
	for i, g := range grid {
		for j, v := range g {
			if v == 1 {
				return islandPerimeterHelper(grid, i, j)
			}
		}
	}
	return 0
}

func islandPerimeterHelper(grid [][]int, x, y int) int {
	// outside
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 {
		return 1
	}
	// water
	if grid[x][y] == 0 {
		return 1
	}
	// visited land
	if grid[x][y] == 2 {
		return 0
	}
	// mark current land
	if grid[x][y] == 1 {
		grid[x][y] = 2
	}
	// dfs
	return islandPerimeterHelper(grid, x+1, y) +
		islandPerimeterHelper(grid, x-1, y) +
		islandPerimeterHelper(grid, x, y+1) +
		islandPerimeterHelper(grid, x, y-1)
}
