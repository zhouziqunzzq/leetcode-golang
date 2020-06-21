package mycode

func removeSliceAt(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func getPermutation(n int, k int) string {
	num := make([]int, n)
	fac := make([]int, n+1)
	ans := make([]rune, n)
	idx := 0

	fac[0] = 1
	for i := 1; i <= n; i++ {
		num[i-1] = i
		fac[i] = fac[i-1] * i
	}

	k--	// because of 0-index
	for i := 1; i <= n; i++ {
		idx = k / fac[n-i]
		ans[i-1] = rune('0' + num[idx])
		num = removeSliceAt(num, idx)
		k -= idx * fac[n-i]
	}

	return string(ans)
}
