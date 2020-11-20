package mycode

// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/solution/
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		cur := nums[m]

		if cur == target {
			return true
		}

		if nums[l] == cur {
			// bin-search won't help, can't decide to go left or right
			l++
			continue
		}

		mf := nums[l] <= cur
		tf := nums[l] <= target
		if mf != tf {
			// mid and target in different sub-arrays
			if mf {
				// mid in first, target in second
				l = m + 1
			} else {
				// target in first, mid in second
				r = m - 1
			}
		} else {
			// mid and target in the same sub-array
			if cur < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return false
}
