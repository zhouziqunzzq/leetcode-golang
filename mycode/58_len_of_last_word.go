package mycode

func lengthOfLastWord(s string) int {
	n := 0
	i := len(s) - 1

	for ; i >= 0 && s[i] == ' '; i-- {
	}

	for ; i >= 0 && s[i] != ' '; i-- {
		n++
	}

	return n
}
