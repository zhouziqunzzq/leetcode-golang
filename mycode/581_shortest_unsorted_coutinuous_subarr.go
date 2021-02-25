package mycode

import "math"

// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/solution/
// Solution 2: brute force adapted from selective sorting, O(n^2)
// Solution 3: sort first and then compare, O(nlogn)
// Solution 4: use stack to determine left and right boundaries, O(n), O(n)
// Solution 5: do without stack, with an extra passes, 4-pass in total, O(n), O(1)
func findUnsortedSubarray(nums []int) int {
	min, max := math.MaxInt32, math.MinInt32
	n := len(nums)

	// 1st pass: find min val after the first descending slope
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			min = minInt(min, nums[i+1])
		}
	}

	// 2nd pass: find max val before the last ascending slope
	for i := n - 1; i > 0; i-- {
		if nums[i-1] > nums[i] {
			max = maxInt(max, nums[i-1])
		}
	}

	l, r := 0, 0
	// 3rd pass: find left boundary
	for i := 0; i < n; i++ {
		if nums[i] > min {
			l = i
			break
		}
	}
	// 4th pass: find right boundary
	for j := n - 1; j >= 0; j-- {
		if nums[j] < max {
			r = j
			break
		}
	}

	if r > l {
		return r - l + 1
	} else {
		return 0
	}
}
