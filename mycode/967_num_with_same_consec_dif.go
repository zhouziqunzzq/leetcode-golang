package mycode

func numsSameConsecDiff(N int, K int) []int {
	ans := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 2; i <= N; i++ {
		newAns := make([]int, 0)
		for _, a := range ans {
			d := a % 10
			if a > 0 && d+K < 10 {
				newAns = append(newAns, a*10+d+K)
			}
			if a > 0 && K > 0 && d-K >= 0 {
				newAns = append(newAns, a*10+d-K)
			}
		}
		ans = newAns
	}

	return ans
}
