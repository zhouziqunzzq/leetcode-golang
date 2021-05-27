package mycode

// https://leetcode.com/problems/maximum-product-subarray/discuss/48230/Possibly-simplest-solution-with-O(n)-time-complexity
func maxProduct152(nums []int) int {
	globalMax := nums[0]
	curMax, curMin := globalMax, globalMax

	for _, n := range nums[1:] {
		if n < 0 {
			lastMax := curMax
			curMax = maxInt(n, curMin*n)
			curMin = minInt(n, lastMax*n)
		} else {
			curMax = maxInt(n, curMax*n)
			curMin = minInt(n, curMin*n)
		}
		globalMax = maxInt(globalMax, curMax)
	}

	return globalMax
}
