package mycode

func countBinarySubstrings(s string) int {
	prevBit, prevCnt := s[0], 0
	curBit, curCnt := s[0], 0
	ans := 0
	i := 0
	n := len(s)

	// init prevCnt
	for i < n && s[i] == prevBit {
		prevCnt++
		i++
	}

	// main loop
	for i < n {
		curBit = s[i]
		curCnt = 0
		for i < n && s[i] == curBit {
			curCnt++
			i++
		}

		ans += minInt(prevCnt, curCnt)

		prevBit, prevCnt = curBit, curCnt
	}

	return ans
}
