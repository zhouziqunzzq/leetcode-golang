package mycode

func minCostToMoveChips(position []int) int {
	odd := 0
	even := 0

	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return minInt(odd, even)
}
