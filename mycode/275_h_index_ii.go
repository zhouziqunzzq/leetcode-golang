package mycode

func hIndex2(citations []int) int {
	N := len(citations)
	l, r := 0, N
	mid := 0
	h := 0
	for l < r {
		mid = (l + r) / 2
		h = N - mid
		if h <= citations[mid] {	// h is too small, increase h by moving left
			r = mid
		} else {	// h is too large, decrease h by moving right
			l = mid + 1
		}
	}

	return N - l
}
