package mycode

func sortedSquares(A []int) []int {
	rst := make([]int, len(A))

	l := 0
	r := len(A) - 1
	p := len(A) - 1

	for l != r && A[l] <= 0 && A[r] >= 0 {
		if A[l]*A[l] >= A[r]*A[r] {
			rst[p] = A[l] * A[l]
			l++
		} else {
			rst[p] = A[r] * A[r]
			r--
		}
		p--
	}

	if A[l] > 0 {
		for l != r {
			rst[p] = A[r] * A[r]
			r--
			p--
		}
	} else if A[r] < 0 {
		for l != r {
			rst[p] = A[l] * A[l]
			l++
			p--
		}
	}
	rst[p] = A[l] * A[l]

	return rst
}
