package mycode

func carPooling(trips [][]int, capacity int) bool {
	stops := make(map[int]int)
	for _, trip := range trips {
		n, s, e := trip[0], trip[1], trip[2]
		if _, ok := stops[s]; !ok {
			stops[s] = n
		} else {
			stops[s] += n
		}
		if _, ok := stops[e]; !ok {
			stops[e] = -n
		} else {
			stops[e] -= n
		}
	}

	cur := 0
	for i := 0; i <= 1000; i++ {
		if n, ok := stops[i]; ok {
			cur += n
		}
		if cur > capacity {
			return false
		}
	}
	return true
}
