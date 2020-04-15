package mycode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	// use ans as from-right-mul arr
	ans[n-1] = 1
	for i := n-2; i >= 0; i-- {
		ans[i] = nums[i+1] * ans[i+1]
	}

	// construct from-left-mul on the fly
	fromLeftMul := 1
	for i := 0; i < n; i++ {
		ans[i] = ans[i] * fromLeftMul
		fromLeftMul *= nums[i]
	}

	return ans
}
