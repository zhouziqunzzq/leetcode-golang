package mycode

func longestMountain(A []int) int {
	if len(A) <= 2 {
		return 0
	}

	// state: 0 - init/flat, 1 - up, 2 - down
	s := 0
	cur := 0
	ans := 0
	prev := A[0]

	for i := 1; i < len(A); i++ {
		diff := A[i] - prev

		if s == 0 {
			if diff > 0 {
				s = 1
				cur = 2
			}
		} else if s == 1 {
			if diff > 0 {
				cur++
			} else if diff < 0 {
				cur++
				s = 2
			} else {
				cur = 0
				s = 0
			}
		} else {
			if diff > 0 {
				cur = 2
				s = 1
			} else if diff == 0 {
				cur = 0
				s = 0
			} else {
				cur++
			}
		}

		if s == 2 {
			ans = maxInt(ans, cur)
		}

		prev = A[i]
	}

	return ans
}
