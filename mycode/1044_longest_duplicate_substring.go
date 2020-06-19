package mycode

// Can be improved by using RK-hash(Rabin-Karp Algorithm)
func longestDupSubstringHelper(S string, substringLen int) string {
	// use a map to check duplicates
	m := make(map[string]bool)
	for i := 0; i < len(S)-substringLen+1; i++ {
		ss := S[i : i+substringLen]
		if _, ok := m[ss]; ok {
			return ss
		} else {
			m[ss] = true
		}
	}

	return ""
}

// Solution 1: Bin-search + map / set / RK-hash  <--- using
// Solution 2: Suffix array  <--- TODO
func longestDupSubstring(S string) string {
	// Hint: Binary search for the length of the answer.
	// (If there's an answer of length 10, then there are answers of length 9, 8, 7, ...)

	l, r := 1, len(S)-1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if longestDupSubstringHelper(S, mid) == "" {
			// mid too large, goto left part
			r = mid - 1
		} else {
			// mid could be larger, goto right part
			l = mid + 1
		}
	}

	return longestDupSubstringHelper(S, l-1)
}
