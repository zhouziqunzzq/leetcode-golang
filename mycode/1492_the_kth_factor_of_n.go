package mycode

func kthFactor(n int, k int) int {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	if k == 1 {
		return n
	} else {
		return -1
	}
}
