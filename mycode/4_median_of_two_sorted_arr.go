package mycode

// https://leetcode.com/problems/median-of-two-sorted-arrays/solution/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	iMin, iMax := 0, m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := iMin + (iMax-iMin)/2
		j := halfLen - i
		if i < iMax && nums1[i] < nums2[j-1] {
			// i is too small
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			// i is too big
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = maxInt(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = minInt(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0
}
