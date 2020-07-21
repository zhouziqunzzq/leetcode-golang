package mycode

func exist(board [][]byte, word string) bool {
	for i, b := range board {
		for j, _ := range b {
			if wordSearchDFS(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordSearchDFS(board [][]byte, word string, x, y, pos int) bool {
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 {
		return false
	}
	if word[pos] != board[x][y] {
		return false
	}
	if pos == len(word)-1 {
		return true
	}

	t := board[x][y]
	board[x][y] = '#'
	rst := wordSearchDFS(board, word, x+1, y, pos+1) ||
		wordSearchDFS(board, word, x-1, y, pos+1) ||
		wordSearchDFS(board, word, x, y+1, pos+1) ||
		wordSearchDFS(board, word, x, y-1, pos+1)
	board[x][y] = t
	return rst
}
