package mycode

func twoSum(nums []int, target int) []int {
	// create index
	index := make(map[int]int)
	for i, v := range nums {
		index[v] = i
	}

	// solve
	for i1, v1 := range nums {
		if i2, ok := index[target-v1]; ok && i1 != i2 {
			if i1 > i2 {
				i1, i2 = i2, i1
			}
			return []int{i1, i2}
		}
	}

	return []int{0, 0}
}
