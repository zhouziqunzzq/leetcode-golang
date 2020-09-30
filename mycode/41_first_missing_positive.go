package mycode

// https://leetcode.com/problems/first-missing-positive/discuss/17071/My-short-c%2B%2B-solution-O(1)-space-and-O(n)-time
func firstMissingPositive(nums []int) int {
	for i, n := range nums {
		for n > 0 && n <= len(nums) && nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n = nums[i]
		}
	}

	for i, n := range nums {
		if n-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}
