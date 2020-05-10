package mycode

func findJudge(N int, trust [][]int) int {
	trustFlag := make([]bool, N+1)
	trusteeCnt := make([]int16, N+1)

	for _, t := range trust {
		trustFlag[t[0]] = true
		trusteeCnt[t[1]]++
	}

	for i := 1; i <= N; i++ {
		if !trustFlag[i] && trusteeCnt[i] == int16(N-1) {
			return i
		}
	}
	return -1
}
