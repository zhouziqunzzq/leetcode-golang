package mycode

func totalNQueens(n int) int {
	return totalNQueensHelper(
		0,
		make(map[int]bool),
		make(map[int]bool),
		make(map[int]bool),
		0,
		n,
	)
}

func totalNQueensHelper(row int, cols, diags, antiDiags map[int]bool, cnt, n int) int {
	for col := 0; col < n; col++ {
		diag := row - col
		antiDiag := row + col
		_, fCol := cols[col]
		_, fDiag := diags[diag]
		_, fAntiDiag := antiDiags[antiDiag]

		if fCol || fDiag || fAntiDiag {
			continue
		}

		if row == n-1 { // base case
			cnt++
		} else {
			cols[col] = true
			diags[diag] = true
			antiDiags[antiDiag] = true

			cnt = totalNQueensHelper(row+1, cols, diags, antiDiags, cnt, n)

			delete(cols, col)
			delete(diags, diag)
			delete(antiDiags, antiDiag)
		}
	}

	return cnt
}
