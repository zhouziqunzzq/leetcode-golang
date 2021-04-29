package mycode

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/discuss/14699/Clean-iterative-solution-with-two-binary-searches-(with-explanation)
// Good explanation about how to move l/r ptrs and set mid in binary search algo.
func searchRange(nums []int, target int) []int {
	rst := make([]int, 2)
	rst[0] = -1
	rst[1] = -1

	if len(nums) == 0 {
		return rst
	}

	// search for left bound
	l, r := 0, len(nums)-1
	mid := 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			// target <= nums[mid]
			r = mid
		}
	}
	if nums[l] != target {
		return rst
	} else {
		rst[0] = l
	}

	// search for right bound
	l, r = 0, len(nums)-1
	for l < r {
		mid = l + (r-l)/2 + 1 // trick to make mid bias to right bound
		if target < nums[mid] {
			r = mid - 1
		} else {
			// nums[mid] <= target
			l = mid
		}
	}
	rst[1] = l

	return rst
}
