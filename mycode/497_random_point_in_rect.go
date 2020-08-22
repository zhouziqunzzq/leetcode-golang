package mycode

import "math/rand"

type Solution struct {
	areaSum     int
	rects       [][]int
	rectEntries []int
}

func Constructor(rects [][]int) Solution {
	s := Solution{areaSum: 0, rects: rects, rectEntries: make([]int, len(rects))}
	for i, r := range rects {
		s.areaSum += (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		s.rectEntries[i] = s.areaSum
	}
	return s
}

func (this *Solution) Pick() []int {
	// first pick a rect randomly with bin-search
	randArea := rand.Int() % (this.areaSum + 1)
	l, r := 0, len(this.rects)-1
	for l < r {
		m := l + (r-l)/2
		if randArea >= this.rectEntries[m] {
			l = m + 1
		} else {
			r = m
		}
	}

	// next pick a point in rect with idx l
	curRect := this.rects[l]
	lb, rb, bb, tb := curRect[0], curRect[2], curRect[1], curRect[3]
	x := lb + rand.Int()%(rb-lb+1)
	y := bb + rand.Int()%(tb-bb+1)
	return []int{x, y}
}

/**
 * Your Solution528 object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
