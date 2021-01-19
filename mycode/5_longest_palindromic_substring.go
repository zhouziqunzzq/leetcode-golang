package mycode

// https://leetcode.com/problems/longest-palindromic-substring/solution/
// Solution 4: expand around center
// caveat: center could be the gap between letters! so 2n-1 centers in total
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// ans is s[maxL:maxR+1]
	maxL, maxR := 0, 0
	for i := 0; i < len(s); i++ {
		// centered at s[i]
		len1 := longestPalindromeHelper(s, i, i)
		// centered between s[i] and s[i+1]
		len2 := longestPalindromeHelper(s, i, i+1)
		maxLen := maxInt(len1, len2)
		if maxLen > maxR-maxL+1 {
			maxL = i - (maxLen-1)/2
			maxR = i + maxLen/2
		}
	}

	return s[maxL : maxR+1]
}

func longestPalindromeHelper(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
