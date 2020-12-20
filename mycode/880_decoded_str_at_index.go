package mycode

// https://leetcode.com/problems/decoded-string-at-index/solution/
func decodeAtIndex(S string, K int) string {
	l := 0 // length of the decoded str
	for _, c := range S {
		if c >= '0' && c <= '9' {
			l *= int(c - '0')
		} else {
			l++
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		c := S[i]
		if c >= '0' && c <= '9' {
			l /= int(c - '0')
			K %= l
		} else {
			if K == l || K == 0 {
				return string(S[i])
			} else {
				l--
			}
		}
	}
	return ""
}
