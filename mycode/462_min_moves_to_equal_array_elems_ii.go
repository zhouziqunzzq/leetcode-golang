package mycode

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1

	ans := 0
	for i < j {
		ans += nums[j] - nums[i]
		i++
		j--
	}
	return ans
}
