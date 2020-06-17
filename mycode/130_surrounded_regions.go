package mycode

func solve130DfsFlag(board [][]byte, x, y int, flags [][]bool) {
	if x < 0 || x >= len(board) ||
		y < 0 || y >= len(board[0]) ||
		flags[x][y] ||
		board[x][y] != 'O' {
		return
	}

	flags[x][y] = true

	solve130DfsFlag(board, x+1, y, flags)
	solve130DfsFlag(board, x-1, y, flags)
	solve130DfsFlag(board, x, y+1, flags)
	solve130DfsFlag(board, x, y-1, flags)
}

func solve130DfsFlip(board [][]byte, x, y int, flags [][]bool) {
	if x < 0 || x >= len(board) ||
		y < 0 || y >= len(board[0]) ||
		flags[x][y] ||
		board[x][y] != 'O' {
		return
	}

	board[x][y] = 'X'
	flags[x][y] = true

	solve130DfsFlip(board, x+1, y, flags)
	solve130DfsFlip(board, x-1, y, flags)
	solve130DfsFlip(board, x, y+1, flags)
	solve130DfsFlip(board, x, y-1, flags)
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	flags := make([][]bool, len(board))
	for i, _ := range flags {
		flags[i] = make([]bool, len(board[0]))
	}

	// flag borders
	for i := 0; i < len(board); i++ {
		solve130DfsFlag(board, i, 0, flags)
		solve130DfsFlag(board, i, len(board[0])-1, flags)
	}
	for j := 0; j < len(board[0]); j++ {
		solve130DfsFlag(board, 0, j, flags)
		solve130DfsFlag(board, len(board)-1, j, flags)
	}

	// flip inner
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			solve130DfsFlip(board, i, j, flags)
		}
	}
}
