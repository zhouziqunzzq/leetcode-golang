package mycode

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/discuss/78993/Clean-and-short-with-comments-C%2B%2B
// Solution: greedy
func increasingTriplet(nums []int) bool {
	c1, c2 := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n <= c1 {
			c1 = n
		} else if n <= c2 {
			c2 = n
		} else {
			// c1 < c2 < n
			// c1(or c1')_index < c2_index < n_index
			return true
		}
	}

	return false
}
