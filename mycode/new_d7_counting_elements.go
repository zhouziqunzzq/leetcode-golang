package mycode

func countElements(arr []int) int {
	elemSet := make(map[int]bool)

	// loop 1: construct set
	for _, elem := range arr {
		elemSet[elem] = true
	}

	// loop 2: calc answer
	ans := 0
	for _, elem := range arr {
		if _, ok := elemSet[elem + 1]; ok {
			ans++
		}
	}

	return ans
}
