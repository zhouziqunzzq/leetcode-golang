package mycode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	opt := make([]int, len(nums))

	// base
	opt[0] = nums[0]
	opt[1] = maxInt(nums[0], nums[1])

	// fill the table
	for i := 2; i < len(nums); i++ {
		opt[i] = maxInt(opt[i-1], nums[i]+opt[i-2])
	}

	return opt[len(opt)-1]
}
