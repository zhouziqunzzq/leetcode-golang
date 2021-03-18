package mycode

// https://leetcode.com/problems/wiggle-subsequence/solution/
// Greedy
// Can also be solved using DP
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	prevDiff := nums[1] - nums[0]
	ans := 2 // beg & end
	if prevDiff == 0 {
		ans = 1 // end only
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			ans++
			prevDiff = diff
		}
	}

	return ans
}
