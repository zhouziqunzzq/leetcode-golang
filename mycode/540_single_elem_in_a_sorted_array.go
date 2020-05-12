package mycode

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		n := m ^ 1
		if nums[m] == nums[n] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}
