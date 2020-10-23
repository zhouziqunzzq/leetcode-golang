package mycode

// https://leetcode.com/problems/132-pattern/solution/
// Approach 4: using a stack to store possible nums[k] in descending order.
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// preprocess
	// min[i] = min(nums[0...i])
	min := make([]int, len(nums))
	min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		min[i] = minInt(min[i-1], nums[i])
	}

	s := make([]int, 0) // s is a stack
	// scan from back to find possible nums[k]
	// 132: nums[i] < nums[k] < nums[j], i < j < k
	// using min[i] to represent 132:
	// min[j] < nums[k] < nums[j], j < k
	for j := len(nums) - 1; j >= 0; j-- {
		// first check if min[j] < nums[j]
		if min[j] < nums[j] {
			// if so, check stack for possible nums[k]
			// we need min[j] < nums[k], so we keep popping if min[j] >= nums[k]
			for len(s) > 0 && min[j] >= s[len(s)-1] {
				s = s[:len(s)-1]
			}
			// now if there's still possible nums[k] in the stack,
			// then we've detected a 132 pattern
			if len(s) > 0 && s[len(s)-1] < nums[j] {
				return true
			}
			// if not, it means that either stack is empty
			// or all possible nums[k] in the stack >= nums[j].
			// in either case, we've found a new possible nums[k]
			// which is current nums[j]. we then push it into stack.
			s = append(s, nums[j])
		}
	}
	return false
}
