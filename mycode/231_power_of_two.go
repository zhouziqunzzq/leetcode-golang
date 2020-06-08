package mycode

func isPowerOfTwo(n int) bool {
	flag := false

	for n > 0 {
		if n&1 == 1 {
			if flag {
				return false
			} else {
				flag = true
			}
		}
		n >>= 1
	}

	return flag
}
