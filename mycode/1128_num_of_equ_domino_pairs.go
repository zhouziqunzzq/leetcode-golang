package mycode

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := make([]int, 100)

	for _, d := range dominoes {
		key := 0
		if d[0] <= d[1] {
			key = 10 * d[0] + d[1]
		} else {
			key = 10 * d[1] + d[0]
		}
		cnt[key]++
	}

	ans := 0
	for _, c := range cnt {
		ans += c * (c-1) / 2
	}

	return ans
}
