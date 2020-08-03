package mycode

func isAlphabetOrNum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func isPalindromeEasy(s string) bool {
	l, r := 0, len(s)-1
	ss := []rune(s)
	for l <= r {
		lc, rc := toLower(ss[l]), toLower(ss[r])
		if isAlphabetOrNum(lc) && isAlphabetOrNum(rc) {
			if lc == rc {
				l++
				r--
				continue
			} else {
				return false
			}
		}
		if !isAlphabetOrNum(lc) {
			l++
		}
		if !isAlphabetOrNum(rc) {
			r--
		}
	}
	return true
}
