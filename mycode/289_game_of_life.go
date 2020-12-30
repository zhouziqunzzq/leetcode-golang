package mycode

// https://leetcode.com/problems/game-of-life/solution/
// Solution 2: O(1) space, two-pass style
func gameOfLife(board [][]int) {
	neighbors := []int{0, 1, -1}
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	// 1st pass: update cells with two intermedia states:
	// -1: died after updated but originally 1
	// 2:  live after updated but originally 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			// count live neighbors
			liveNeighbors := 0
			for _, dr := range neighbors {
				for _, dc := range neighbors {
					// skip the cell itself
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if (nr >= 0 && nr < m) && (nc >= 0 && nc < n) &&
						(board[nr][nc] == 1 || board[nr][nc] == -1) {
						liveNeighbors++
					}
				}
			}

			// rule 1 & 3
			if board[r][c] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[r][c] = -1
			}
			// rule 4
			if board[r][c] == 0 && liveNeighbors == 3 {
				board[r][c] = 2
			}
		}
	}

	// 2nd pass: recover 0-1 cell states
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] > 0 {
				board[r][c] = 1
			} else {
				board[r][c] = 0
			}
		}
	}
}
