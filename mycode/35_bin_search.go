package mycode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 0
	r := len(nums) - 1
	m := 0

	for l < r {
		m = (l+r)/2
		if nums[m] == target {
			return m
		} else if target < nums[m] {
			r = m-1
		} else {
			l = m+1
		}
	}

	if target > nums[l] {
		return l+1
	} else {
		return l
	}
}
