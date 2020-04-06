package mycode

import "sort"

func minIncrementForUnique(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	ans := 0
	//tmp := make([]int, len(A))
	sort.Ints(A)
	t := A[0]
	//tmp[0] = t
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			t++
			ans += t - A[i]
		} else {
			if A[i] <= t {
				t++
				ans += t - A[i]
			} else {
				t = A[i]
			}
		}
	}
	return ans
}
