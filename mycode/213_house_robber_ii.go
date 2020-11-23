package mycode

func rob2(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return maxInt(nums[0], nums[1])
	}

	// max reward for robbing [0...i] where i <= n-2
	opt1 := make([]int, n)
	// max reward for robbing [1...i] where i <= n-1
	opt2 := make([]int, n)

	opt1[0] = nums[0]
	opt1[1] = maxInt(nums[0], nums[1])
	opt2[1] = nums[1]
	opt2[2] = maxInt(nums[1], nums[2])

	for i := 2; i < n; i++ {
		if i >= 2 && i <= n-2 {
			opt1[i] = maxInt(opt1[i-1], nums[i]+opt1[i-2])
		}
		if i >= 3 && i <= n-1 {
			opt2[i] = maxInt(opt2[i-1], nums[i]+opt2[i-2])
		}
	}

	return maxInt(opt1[n-2], opt2[n-1])
}
