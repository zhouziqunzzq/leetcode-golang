package mycode

func runningSum(nums []int) []int {
	rst := make([]int, len(nums))
	rst[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		rst[i] = rst[i-1] + nums[i]
	}

	return rst
}
