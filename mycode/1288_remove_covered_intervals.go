package mycode

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		il, ir, jl, jr := intervals[i][0], intervals[i][1], intervals[j][0], intervals[j][1]
		if il < jl {
			return true
		} else if il == jl {
			return (ir - il) > (jr - jl)
		} else {
			return false
		}
	})

	ans := 0
	cur := 0
	i := 0
	for i < len(intervals) {
		cur = intervals[i][1]
		ans++
		i++
		for i < len(intervals) && intervals[i][1] <= cur {
			i++
		}
	}

	return ans
}
