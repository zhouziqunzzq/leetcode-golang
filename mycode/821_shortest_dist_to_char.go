package mycode

import "math"

func shortestToChar(s string, c byte) []int {
	rst := make([]int, len(s))
	prev := math.MinInt32
	cur := math.MinInt32

	for i := range s {
		if i > cur {
			// move prev and cur
			prev = cur
			for cur = i; cur < len(s) && s[cur] != c; cur++ {
			}
			if cur == len(s) {
				cur = math.MaxInt32
			}
		}
		rst[i] = minInt(i-prev, cur-i)
	}

	return rst
}
