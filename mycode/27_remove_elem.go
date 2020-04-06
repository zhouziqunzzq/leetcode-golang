package mycode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	} else {
		l := 0
		r := len(nums) - 1

		for l != r {
			if nums[l] == val {
				for l != r && nums[r] == val {
					r--
				}
				if l == r {
					return r
				} else {
					nums[l], nums[r] = nums[r], nums[l]
					r--
				}
			} else {
				l++
			}
		}

		if nums[l] == val {
			return l
		} else {
			return l+1
		}
	}
}
