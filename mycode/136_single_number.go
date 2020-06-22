package mycode

func singleNumberOne(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}

	return ans
}
