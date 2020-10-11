package mycode

// https://leetcode.com/problems/remove-duplicate-letters/discuss/76768/A-short-O(n)-recursive-greedy-solution
func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}

	var cnt [26]int
	ss := []rune(s)
	for _, c := range ss {
		cnt[c-'a']++
	}

	// greedily choose cur such that:
	// 1. s[cur:] contains all letters
	// 2. s[cur] is as small as possible
	// 3. cur is as small as possible
	cur := 0
	for i, c := range ss {
		if c < ss[cur] {
			cur = i
		}
		cnt[c-'a']--
		if cnt[c-'a'] == 0 {
			break
		}
	}

	// construct sub-problem
	nextSs := make([]rune, 0, len(ss)-cur-1)
	for i := cur + 1; i < len(ss); i++ {
		if ss[i] != ss[cur] {
			nextSs = append(nextSs, ss[i])
		}
	}

	// recursively solve the sub-problem
	return string(ss[cur]) + removeDuplicateLetters(string(nextSs))
}
