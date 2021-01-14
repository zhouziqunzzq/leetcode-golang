package mycode

import "math"

// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/discuss/935935/Java-Detailed-Explanation-O(N)-Prefix-SumMap-Longest-Target-Sub-Array
// reduce to finding the maximum subarray summing up to sum(nums)-x
func minOperations(nums []int, x int) int {
	t := -x
	for _, n := range nums {
		t += n
	}
	if t == 0 {
		return len(nums)
	}

	preSum := make(map[int]int)
	preSum[0] = -1
	sum := 0
	maxLen := math.MinInt32

	for i, n := range nums {
		sum += n
		if prevIdx, ok := preSum[sum-t]; ok {
			maxLen = maxInt(maxLen, i-prevIdx)
		}

		preSum[sum] = i
	}

	if maxLen == math.MinInt32 {
		return -1
	} else {
		return len(nums) - maxLen
	}
}
