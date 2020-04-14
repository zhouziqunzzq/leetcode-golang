package mycode

func stringShift(s string, shift [][]int) string {
	finalShift := 0

	for _, s := range shift {
		if s[0] == 0 {
			// shift left
			finalShift -= s[1]
		} else {
			// shift right
			finalShift += s[1]
		}
	}

	if finalShift == 0 {
		return s
	} else if finalShift < 0 {
		// shift left
		finalShift = 0 - finalShift
		finalShift %= len(s)
		return s[finalShift:] + s[:finalShift]
	} else {
		// shift right
		finalShift %= len(s)
		return s[len(s)-finalShift:] + s[:len(s)-finalShift]
	}
}
