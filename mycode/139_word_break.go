package mycode

import "strings"

// https://leetcode.com/problems/word-break/discuss/43790/Java-implementation-using-DP-in-two-ways
func wordBreak(s string, wordDict []string) bool {
	opt := make([]bool, len(s)+1)
	opt[0] = true

	for i := 0; i < len(s); i++ {
		if !opt[i] {
			continue
		}
		for _, word := range wordDict {
			if strings.HasPrefix(s[i:], word) {
				opt[i+len(word)] = true
			}
		}
	}

	return opt[len(s)]
}
