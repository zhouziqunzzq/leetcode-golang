package mycode

func kLengthApart(nums []int, k int) bool {
	prev := -len(nums) - 1

	for i, v := range nums {
		if v == 1 {
			if i-prev-1 < k {
				return false
			} else {
				prev = i
			}
		}
	}

	return true
}
