package mycode

func intervalIntersection(A [][]int, B [][]int) [][]int {
	pa, pb := 0, 0
	ans := make([][]int, 0)

	for pa < len(A) && pb < len(B) {
		l := maxInt(A[pa][0], B[pb][0])
		r := minInt(A[pa][1], B[pb][1])
		if l <= r {
			ans = append(ans, make([]int, 2))
			ans[len(ans)-1][0] = l
			ans[len(ans)-1][1] = r
		}
		if r == A[pa][1] {
			pa++
		} else {
			pb++
		}
	}

	return ans
}
