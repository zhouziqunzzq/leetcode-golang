package mycode

func reverseString(s []byte) {
	if s == nil || len(s) <= 1 {
		return
	} else if len(s) == 2 {
		s[0], s[1] = s[1], s[0]
		return
	} else {
		s[0], s[len(s)-1] = s[len(s)-1], s[0]
		reverseString(s[1 : len(s)-1])
	}
}
