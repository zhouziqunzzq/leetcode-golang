package mycode

func distributeCandies(candyType []int) int {
	n := len(candyType)
	cnt := make(map[int]bool)

	for _, c := range candyType {
		if _, ok := cnt[c]; !ok {
			cnt[c] = true
		}
	}

	return minInt(len(cnt), n/2)
}
