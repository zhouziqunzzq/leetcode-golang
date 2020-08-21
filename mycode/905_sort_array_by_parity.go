package mycode

func sortArrayByParity(A []int) []int {
	pe, po := 0, len(A)-1
	for pe < po {
		if A[pe]%2 == 0 {
			pe++
		} else {
			A[pe], A[po] = A[po], A[pe]
			po--
		}
	}
	return A
}
