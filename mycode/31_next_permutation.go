package mycode

// https://leetcode.com/problems/next-permutation/solution/
func nextPermutation(nums []int) {
	i := len(nums) - 2

	// find the leftmost i s.t. nums[i] < nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// find next num for the i-th place
		// which is the rightmost j s.t.
		// j in [i+1...N) and nums[j] > nums[i]
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	reverseSlice(nums[i+1:])
}
