package mycode

// https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/60294/Solution-explained
// Solution: Quick selection, avg O(n)
func findKthLargest(nums []int, k int) int {
	k = len(nums) - k // reduce k-th largest to the elem ranking len(nums)-k
	l, r := 0, len(nums)-1

	for l < r {
		pivotIdx := findKthLargestPartition(nums, l, r)
		if pivotIdx < k {
			l = pivotIdx + 1
		} else if pivotIdx > k {
			r = pivotIdx - 1
		} else {
			break
		}
	}

	return nums[k]
}

// Lomuto partition algorithm
// use nums[r] as pivot to partition nums[l...r] into two parts
func findKthLargestPartition(nums []int, l, r int) int {
	leftPartP := l
	pivot := nums[r]

	for i := l; i < r; i++ {
		if nums[i] <= pivot {
			nums[i], nums[leftPartP] = nums[leftPartP], nums[i]
			leftPartP++
		}
	}

	nums[r], nums[leftPartP] = nums[leftPartP], nums[r]

	return leftPartP
}
