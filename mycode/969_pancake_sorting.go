package mycode

func pancakeSort(A []int) []int {
	rst := make([]int, 0)

	for x := len(A); x >= 1; x-- {
		if A[x-1] == x {
			continue
		}
		i := 0
		for i = 0; A[i] != x && i < len(A); i++ {
		}
		pancakeSortReverse(A, i+1)
		rst = append(rst, i+1)
		pancakeSortReverse(A, x)
		rst = append(rst, x)
	}

	return rst
}

// reverse A[0...k-1]
func pancakeSortReverse(A []int, k int) {
	l, r := 0, k-1
	for l < r {
		A[l], A[r] = A[r], A[l]
		l++
		r--
	}
}
