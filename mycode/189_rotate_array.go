package mycode

// 1 2 3 4 5 6 7
func rotate189(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}

	processRight := true
	t := k
	if t > n/2 {
		t = n - k
		processRight = false
	}

	l, r := 0, n-1
	for i := 0; i < t; i++ {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	rotateHelperReverse(nums, l, r)
	if processRight {
		rotateHelperReverse(nums, 0, l-1)
		rotateHelperReverse(nums, l, n-1)
	} else {
		rotateHelperReverse(nums, r+1, n-1)
		rotateHelperReverse(nums, 0, r)
	}
}

// 0 1 2 3 4 5 6
func rotateHelperReverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
