package mycode

import "math"

// https://leetcode.com/problems/maximum-gap/discuss/50643/bucket-sort-JAVA-solution-with-explanation-O(N)-time-and-space
// Bucket sort!
func maximumGap(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	n := len(nums)

	min, max := nums[0], nums[0]
	for _, v := range nums {
		min = minInt(min, v)
		max = maxInt(max, v)
	}

	gap := int(math.Ceil(float64(max-min) / float64(n-1)))
	bucketMin := make([]int, n-1)
	bucketMax := make([]int, n-1)
	fillIntSlice(bucketMin, math.MaxInt32)
	fillIntSlice(bucketMax, math.MinInt32)

	for _, v := range nums {
		if v == min || v == max {
			continue
		}
		idx := (v - min) / gap
		bucketMin[idx] = minInt(bucketMin[idx], v)
		bucketMax[idx] = maxInt(bucketMax[idx], v)
	}

	maxGap := math.MinInt32
	prev := min
	for i := range bucketMin {
		if bucketMin[i] == math.MaxInt32 || bucketMax[i] == math.MinInt32 { // empty bucket
			continue
		}
		maxGap = maxInt(maxGap, bucketMin[i]-prev)
		prev = bucketMax[i]
	}
	maxGap = maxInt(maxGap, max-prev) // don't forget the final gap

	return maxGap
}
