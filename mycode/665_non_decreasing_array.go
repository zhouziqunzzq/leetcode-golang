package mycode

// https://leetcode.com/problems/non-decreasing-array/discuss/106826/JavaC%2B%2B-Simple-greedy-like-solution-with-explanation
// Greedy: prefer to change nums[i-1] than nums[i]
func checkPossibility(nums []int) bool {
	cnt := 0

	for i := 1; i < len(nums) && cnt <= 1; i++ {
		if nums[i] < nums[i-1] {
			cnt++
			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i] // greedy: prefer this fix...
			} else {
				nums[i] = nums[i-1] // ...than this one
			}
		}
	}

	return cnt <= 1
}
