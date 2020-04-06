package mycode

func isHappy(n int) bool {
	sum := 0
	for n != 1 && n != 4 {
		sum = 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}

	return n == 1
}
