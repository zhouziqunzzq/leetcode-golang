package mycode

func maxPower(s string) int {
	ans := 1
	cur := 1
	curChar := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == curChar {
			cur++
			ans = maxInt(ans, cur)
		} else {
			curChar = s[i]
			cur = 1
		}
	}

	return ans
}
