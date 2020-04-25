package mycode

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	opt := make([]bool, len(nums))

	// base case
	opt[0] = nums[0] > 0

	// fill the table
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if opt[j] && j+nums[j] >= i {
				opt[i] = true
				break
			}
		}
	}

	return opt[len(opt)-1]
}

func canJumpGreedy(nums []int) bool {
	lastPos := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
