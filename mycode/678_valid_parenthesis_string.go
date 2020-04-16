package mycode

func checkValidString(s string) bool {
	// [lo, hi] forms the range of possible balance value
	lo, hi := 0, 0

	for _, c := range s {
		if c == '(' {
			lo++
			hi++
		} else if c == ')' {
			lo--
			hi--
		} else {	// c == '*'
			lo--
			hi++
		}
		if hi < 0 {
			return false
		}
		if lo < 0 {
			lo = 0
		}
	}

	return lo == 0
}
