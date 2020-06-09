package mycode

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sp := 0
	ss := []rune(s)

	for _, c := range t {
		if ss[sp] == c {
			sp++
		}
		if sp == len(s) {
			return true
		}
	}

	return false
}
