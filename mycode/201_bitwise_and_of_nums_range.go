package mycode

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}

	t := 0x80000000
	ans := 0

	for m & t == n & t && t > 0 {
		ans |= m & t
		t >>= 1
	}

	return ans
}
