package mycode

// Solution 1: reduce products to sums by taking log, then use bin-search on prefix sum array
// Solution 2: sliding window (two pointers)  <------ implement this
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	ans := 0
	l := 0
	prod := 1
	for r, n := range nums {
		prod *= n
		for prod >= k && l <= r {
			prod /= nums[l]
			l++
		}
		ans += r - l + 1 // this is the tricky part
	}

	return ans
}
