package mycode

import "fmt"

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)

	if len(nums) == 0 {
		return ans
	}

	l, r := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			r = nums[i]
		} else {
			if l == r {
				ans = append(ans, fmt.Sprintf("%d", l))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", l, r))
			}
			l = nums[i]
			r = nums[i]
		}
	}
	if l == r {
		ans = append(ans, fmt.Sprintf("%d", l))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", l, r))
	}
	return ans
}
