package mycode

import "math"

// https://leetcode.com/problems/minimize-deviation-in-array/discuss/955262/C%2B%2B-intuitions-and-flip
func minimumDeviation(nums []int) int {
	curMin := math.MaxInt32
	pq := CreateIntPriorityQueue(true)

	// multiple all odd nums by two, and then we can
	// focus purely on the divide operation
	for _, n := range nums {
		if n%2 != 0 {
			n *= 2
		}
		curMin = minInt(curMin, n)
		pq.PushInt(n)
	}

	ans := math.MaxInt32
	for pq.PeekInt()%2 == 0 {
		ans = minInt(ans, pq.PeekInt()-curMin)
		curMin = minInt(curMin, pq.PeekInt()/2)
		pq.PushInt(pq.PeekInt() / 2)
		_ = pq.PopInt()
	}

	return minInt(ans, pq.PeekInt()-curMin)
}
