package mycode

func majorityElement169(nums []int) int {
	n := len(nums)
	cntMap := make(map[int]int)

	for _, v := range nums {
		if _, ok := cntMap[v]; ok {
			cntMap[v]++
			if cntMap[v] > (n / 2) {
				return v
			}
		} else {
			cntMap[v] = 1
		}
	}

	for _, v := range nums {
		if cnt, ok := cntMap[v]; ok && cnt > (n / 2) {
			return v
		}
	}
	return 0
}
