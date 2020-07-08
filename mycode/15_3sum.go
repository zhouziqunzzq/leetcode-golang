package mycode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] { // skip duplicates
			continue
		}
		l := i + 1
		r := len(nums) - 1
		// we want nums[l] + nums[r] + v == 0
		for l < r {
			if nums[l]+nums[r]+v > 0 { // too large
				r--
			} else if nums[l]+nums[r]+v < 0 { // too small
				l++
			} else { // find a feasible tuple
				t := make([]int, 3)
				t[0], t[1], t[2] = v, nums[l], nums[r]
				ans = append(ans, t)
				// skip duplicates
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				r--
				l++
			}
		}
	}

	return ans
}
