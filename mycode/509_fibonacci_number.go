package mycode

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fp2, fp1 := 0, 1
	ans := fp2 + fp1
	n -= 2

	for n > 0 {
		fp2, fp1 = fp1, ans
		ans = fp2 + fp1
		n--
	}

	return ans
}
