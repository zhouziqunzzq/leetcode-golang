package mycode

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	a := make([]int, n+1)
	a[1] = 1
	ans := 1

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			a[i] = a[i/2]
		} else {
			a[i] = a[i/2] + a[(i+1)/2]
		}
		ans = maxInt(ans, a[i])
	}

	return ans
}
