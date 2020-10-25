package mycode

// https://leetcode.com/problems/stone-game-iv/solution/
// Nice problem with both memo-dfs or DP solution :p
// Let's try memo-dfs here.
func winnerSquareGame(n int) bool {
	memo := make(map[int]bool)

	// base cases
	memo[0] = false
	memo[1] = true

	return winnerSquareGameDFS(n, memo)
}

func winnerSquareGameDFS(n int, memo map[int]bool) bool {
	if rst, ok := memo[n]; ok {
		return rst
	}

	canWin := false
	for i := 1; i*i <= n; i++ {
		// try to push adversary into a losing state
		if winnerSquareGameDFS(n-i*i, memo) == false {
			canWin = true
			break
		}
	}

	memo[n] = canWin
	return canWin
}
