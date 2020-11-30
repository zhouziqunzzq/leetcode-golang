package mycode

import "sort"

// https://briangordon.github.io/2014/08/the-skyline-problem.html
func getSkyline(buildings [][]int) [][]int {
	// sort critical points
	cp := make([][]int, len(buildings)*2)
	for i, b := range buildings {
		// start of a building, height stored as negative
		cp[2*i] = make([]int, 2)
		cp[2*i][0] = b[0]
		cp[2*i][1] = -b[2]
		// end of a building, height stored as positive
		cp[2*i+1] = make([]int, 2)
		cp[2*i+1][0] = b[1]
		cp[2*i+1][1] = b[2]
	}
	sort.Slice(cp, func(i, j int) bool {
		if cp[i][0] == cp[j][0] {
			return cp[i][1] < cp[j][1]
		} else {
			return cp[i][0] < cp[j][0]
		}
	})

	rst := make([][]int, 0)
	// a priority queue to store heights of active critical points
	pq := CreateIntPriorityQueue(true)
	pq.PushInt(0)
	prevHeight := 0

	// scan through critical points
	for _, p := range cp {
		x, h := p[0], p[1]
		if h < 0 {
			// start point, add height to pq
			pq.PushInt(-h)
		} else {
			// end point, remove height from pq
			pq.RemoveInt(h)
		}

		curMaxHeight := pq.PeekInt()
		if curMaxHeight != prevHeight {
			t := make([]int, 2)
			t[0] = x
			t[1] = curMaxHeight
			rst = append(rst, t)
			prevHeight = curMaxHeight
		}
	}

	return rst
}
