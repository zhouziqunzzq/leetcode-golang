package mycode

func solveNQueens(n int) [][]string {
	state := make([][]rune, n)
	for i := range state {
		state[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			state[i][j] = '.'
		}
	}
	ans := make([][]string, 0)

	solveNQueensHelper(
		state,
		0,
		make(map[int]bool),
		make(map[int]bool),
		make(map[int]bool),
		&ans,
	)

	return ans
}

func solveNQueensHelper(state [][]rune, row int, cols, diags, antiDiags map[int]bool, ans *[][]string) {
	n := len(state)

	if row == n {
		t := make([]string, n)
		for i, r := range state {
			for _, v := range r {
				t[i] += string(v)
			}
		}
		*ans = append(*ans, t)
		return
	}

	for col := 0; col < n; col++ {
		diag := row - col
		antiDiag := row + col
		_, fCol := cols[col]
		_, fDiag := diags[diag]
		_, fAntiDiag := antiDiags[antiDiag]

		if fCol || fDiag || fAntiDiag {
			continue
		}

		cols[col] = true
		diags[diag] = true
		antiDiags[antiDiag] = true
		state[row][col] = 'Q'

		solveNQueensHelper(state, row+1, cols, diags, antiDiags, ans)

		delete(cols, col)
		delete(diags, diag)
		delete(antiDiags, antiDiag)
		state[row][col] = '.'
	}
}
