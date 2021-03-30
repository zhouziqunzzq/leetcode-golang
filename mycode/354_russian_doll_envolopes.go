package mycode

import "sort"

// https://leetcode.com/problems/russian-doll-envelopes/discuss/82763/Java-NLogN-Solution-with-Explanation
// This is essentially a 2-d LIS(Longest Increasing Subsequence) problem.
// Sorting on one dimension reduce the 2-d problem to 1-d.
// Then we can do a standard LIS algorithm in O(nlogn).
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) < 2 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0] // ascending in width(1st dimension)
		} else {
			return envelopes[i][1] > envelopes[j][1] // descending in height(2nd dimension)
		}
	})

	dp := make([]int, len(envelopes))
	size := 0
	for _, e := range envelopes {
		idx := intSlicesBinarySearch(dp, 0, size, e[1])
		dp[idx] = e[1]
		if idx == size { // we've found a longer increasing subsequence
			size++
		}
	}

	return size
}
