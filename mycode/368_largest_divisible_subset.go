package mycode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	opt := make([]int, len(nums))
	optLast := make([]int, len(nums))
	globalMax := 0
	globalIdx := -1

	for i := 0; i < len(nums); i++ {
		opt[i] = 1
		optLast[i] = -1
		// opt[i] = max(opt[j] for j = 0 ... i-1)
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && opt[i] < opt[j]+1 {
				opt[i] = opt[j] + 1
				optLast[i] = j
			}
		}
		if globalMax < opt[i] {
			globalMax = opt[i]
			globalIdx = i
		}
	}

	ans := make([]int, 0)
	for globalIdx != -1 {
		ans = append(ans, nums[globalIdx])
		globalIdx = optLast[globalIdx]
	}

	return ans
}
