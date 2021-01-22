package mycode

import "sort"

// https://leetcode.com/problems/determine-if-two-strings-are-close/discuss/935920/C%2B%2B-Short-and-Simple-oror-O(-N-)-solution
func closeStrings(word1 string, word2 string) bool {
	var cnt1, cnt2 [26]int
	var occ1, occ2 [26]bool

	for _, c := range word1 {
		cnt1[c-'a']++
		occ1[c-'a'] = true
	}

	for _, c := range word2 {
		cnt2[c-'a']++
		occ2[c-'a'] = true
	}

	sort.Ints(cnt1[:])
	sort.Ints(cnt2[:])

	return cnt1 == cnt2 && occ1 == occ2
}
