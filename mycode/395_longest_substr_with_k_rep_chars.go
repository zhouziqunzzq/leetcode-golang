package mycode

// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/solution/
// Solution 1: divide and conquer
// Solution 2: Sliding window <--
func longestSubstring(s string, k int) int {
	rst := 0
	countMap := make([]int, 26)

	// compute max number of unique chars in s
	maxUnique := 0
	for _, c := range s {
		if countMap[c-'a'] == 0 {
			maxUnique++
		}
		countMap[c-'a']++
	}

	// try different number of unique chars from 1 to maxUnique
	for curUnique := 1; curUnique <= maxUnique; curUnique++ {
		countMap = make([]int, 26)
		l, r := 0, 0
		unique := 0      // num of unique chars in the sliding window [l...r)
		cntAtLeastK := 0 // num of unique chars with counting >= K in the sliding window

		for r < len(s) {
			if unique <= curUnique {
				// expand window to the right
				idx := s[r] - 'a'
				if countMap[idx] == 0 {
					unique++
				}
				countMap[idx]++
				if countMap[idx] == k {
					cntAtLeastK++
				}
				r++
			} else {
				// shrink window from the left
				idx := s[l] - 'a'
				if countMap[idx] == k {
					cntAtLeastK--
				}
				countMap[idx]--
				if countMap[idx] == 0 {
					unique--
				}
				l++
			}

			// update global maximum if possible
			if unique == curUnique && unique == cntAtLeastK {
				rst = maxInt(rst, r-l)
			}
		}
	}

	return rst
}
