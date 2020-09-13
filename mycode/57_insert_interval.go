package mycode

func insert(intervals [][]int, newInterval []int) [][]int {
	l, r := 0, 0
	crossL, crossR := false, false

	for l = 0; l < len(intervals) && newInterval[0] >= intervals[l][0]; l++ {
		if newInterval[0] >= intervals[l][0] && newInterval[0] <= intervals[l][1] {
			crossL = true
			break
		}
	}

	for r = l; r < len(intervals) && newInterval[1] >= intervals[r][0]; r++ {
		if newInterval[1] >= intervals[r][0] && newInterval[1] <= intervals[r][1] {
			crossR = true
			break
		}
	}

	// left intervals
	rst := make([][]int, l)
	i := 0
	for ; i < l; i++ {
		rst[i] = make([]int, 2)
		rst[i][0] = intervals[i][0]
		rst[i][1] = intervals[i][1]
	}

	// merged intervals
	ml := newInterval[0]
	if crossL {
		ml = intervals[l][0]
	}
	mr := newInterval[1]
	if crossR {
		mr = intervals[r][1]
	}
	ti := make([]int, 2)
	ti[0] = ml
	ti[1] = mr
	rst = append(rst, ti)

	// right intervals
	i = r
	if crossR {
		i = r + 1
	}
	for ; i < len(intervals); i++ {
		ti = make([]int, 2)
		ti[0] = intervals[i][0]
		ti[1] = intervals[i][1]
		rst = append(rst, ti)
	}

	return rst
}
