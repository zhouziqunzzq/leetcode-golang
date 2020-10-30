package mycode

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	lengths := make([]int, n)
	counts := make([]int, n)
	for i := range counts {
		counts[i] = 1
	}

	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				if lengths[i] >= lengths[j] {
					lengths[j] = lengths[i] + 1
					counts[j] = counts[i]
				} else if lengths[i]+1 == lengths[j] {
					counts[j] += counts[i]
				}
			}
		}
	}

	longest, ans := 0, 0
	for _, l := range lengths {
		longest = maxInt(longest, l)
	}
	for i, l := range lengths {
		if l == longest {
			ans += counts[i]
		}
	}

	return ans
}
