package mycode

import "math"

func maxSubarraySumCircularHelper(A []int, isRevert bool) int {
	cur := 0
	ans := math.MinInt32
	for _, v := range A {
		if isRevert {
			v = -v
		}
		if cur > 0 {
			cur += v
		} else {
			cur = v
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func maxSubarraySumCircular(A []int) int {
	if len(A) == 1 {
		return A[0]
	}

	S := 0
	for _, v := range A {
		S += v
	}
	ans1 := maxSubarraySumCircularHelper(A, false)
	ans2 := S + maxSubarraySumCircularHelper(A[1:len(A)-1], true)

	if ans1 > ans2 {
		return ans1
	} else {
		return ans2
	}
}
