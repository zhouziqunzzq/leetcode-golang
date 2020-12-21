package mycode

import "sort"

// https://leetcode.com/problems/smallest-range-ii/solution/
// Greedy
func smallestRangeII(A []int, K int) int {
	n := len(A)
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	rst := A[n-1] - A[0]
	t1, t2 := A[0]+K, A[n-1]-K
	for i := 0; i < n-1; i++ {
		// greedy: if A[i] < A[j], then A[i]+K, A[j]-K is better than
		// A[i]-K, A[j]+K.
		low := minInt(t1, A[i+1]-K)
		high := maxInt(t2, A[i]+K)
		rst = minInt(rst, high-low)
	}

	return rst
}
