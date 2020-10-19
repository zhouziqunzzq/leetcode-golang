package mycode

func minDominoRotations(A []int, B []int) int {
	if len(A) != len(B) {
		return -1
	}
	n := len(A)
	if n == 0 {
		return 0
	}

	ans := -1
	x, y := A, B

	for k := 0; k < 2; k++ {
		for t := 1; t <= 6; t++ {
			cur := 0
			isSuccess := true
			for i := 0; i < n; i++ {
				// try to make x[i] = t for all i
				if x[i] != t {
					if y[i] == t {
						cur++
					} else {
						isSuccess = false
						break
					}
				}
			}
			if isSuccess {
				if ans == -1 || cur < ans {
					ans = cur
				}
			}
		}
		x, y = y, x
	}

	return ans
}
