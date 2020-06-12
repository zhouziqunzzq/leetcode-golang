package mycode

import (
	"math/rand"
	"time"
)

type Solution struct {
	accW []int // accumulated weights
	sumW int
}

func Constructor528(w []int) Solution {
	rand.Seed(time.Now().UnixNano())

	s := Solution{
		accW: make([]int, len(w)),
	}
	s.accW[0] = w[0]
	s.sumW += w[0]
	for i := 1; i < len(w); i++ {
		s.accW[i] = s.accW[i-1] + w[i]
		s.sumW += w[i]
	}

	return s
}

func (this *Solution) PickIndex() int {
	randAccW := rand.Intn(this.sumW) + 1 // [1, sumW]

	// perform bin-search
	l, r := 0, len(this.accW)-1
	for l <= r {
		mid := (l + r) / 2
		if this.accW[mid] == randAccW {
			return mid
		} else if randAccW < this.accW[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
