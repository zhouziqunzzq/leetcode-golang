package mycode

func missingNumber(nums []int) int {
	n := len(nums)
	rst := 0

	for _, v := range nums {
		rst ^= v
	}

	for i := 0; i <= n; i++ {
		rst ^= i
	}

	return rst
}
