package mycode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	p := 0

	for i := 0; i < len(haystack); {
		if haystack[i] == needle[p] {
			i++
			p++
			if p == len(needle) {
				return i - len(needle)
			}
		} else {
			i -= p-1
			p = 0
		}
	}

	return -1
}
