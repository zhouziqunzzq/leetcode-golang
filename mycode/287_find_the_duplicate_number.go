package mycode

// Approach 3: Floyd's Tortoise and Hare (Cycle Detection)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// phase one
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// phase two
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
