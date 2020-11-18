package mycode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l, r := 0, 0
	ans := make([][]int, 0)

	for i := 0; i < len(intervals); {
		l, r = intervals[i][0], intervals[i][1]
		i++

		for i < len(intervals) && intervals[i][0] <= r {
			r = maxInt(r, intervals[i][1])
			i++
		}

		t := make([]int, 2)
		t[0], t[1] = l, r
		ans = append(ans, t)
	}

	return ans
}
