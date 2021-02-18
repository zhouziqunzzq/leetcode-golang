package mycode

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	ans := 0
	l, r := 0, 1
	d := A[1] - A[0]
	for r < n {
		if (r < n-1) && A[r+1]-A[r] == d {
			// extend slice
			r++
		} else {
			// update ans
			sliceLen := r - l + 1
			if sliceLen >= 3 {
				ans += (sliceLen-2)*sliceLen - ((sliceLen + 1) * (sliceLen - 2) / 2)
			}

			// reset slice
			l = r
			r++
			if r < n {
				d = A[r] - A[l]
			}
		}
	}

	return ans
}

// also, see recursive and DP solution at: https://leetcode.com/problems/arithmetic-slices/solution/
