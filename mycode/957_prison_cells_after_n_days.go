package mycode

func prisonAfterNDays(cells []int, N int) []int {
	stateToIdx := make(map[int]int)
	cycleLen := 0
	remainLen := 0

	for i := 0; i < N; i++ {
		state := prisonAfterNDaysToInt(cells)

		if idx, ok := stateToIdx[state]; ok {
			// cycle found
			cycleLen = i - idx + 1
			remainLen = (N - i) % (cycleLen - 1)
			break
		}

		stateToIdx[state] = i
		cells = prisonAfterNDaysGetNext(cells)
	}

	for i := 0; i < remainLen; i++ {
		cells = prisonAfterNDaysGetNext(cells)
	}

	return cells
}

func prisonAfterNDaysGetNext(cells []int) []int {
	newCells := make([]int, len(cells))
	for i := 1; i < len(cells)-1; i++ {
		if cells[i-1] == cells[i+1] {
			newCells[i] = 1
		} else {
			newCells[i] = 0
		}
	}
	return newCells
}

func prisonAfterNDaysToInt(cells []int) int {
	rst := 0
	for i, c := range cells {
		rst |= c << i
	}
	return rst
}
