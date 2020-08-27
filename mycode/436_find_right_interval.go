package mycode

import "sort"

type findRightIntervalStruct struct {
	start int
	idx   int
}

func findRightInterval(intervals [][]int) []int {
	rst := make([]int, len(intervals))
	a := make([]findRightIntervalStruct, len(intervals))

	for i, interval := range intervals {
		a[i].start = interval[0]
		a[i].idx = i
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].start < a[j].start
	})

	for i, interval := range intervals {
		// bin-search for lower bound
		t := interval[1]
		l, r := 0, len(a)-1
		for l <= r && l < len(a) {
			m := l + (r-l)/2
			mt := a[m].start
			if mt < t {
				l = m + 1
			} else {
				// mt >= t
				r = m - 1
			}
		}
		if l >= len(a) {
			rst[i] = -1
		} else {
			rst[i] = a[l].idx
		}
	}

	return rst
}
