package mycode

func powerfulIntegers(x int, y int, bound int) []int {
	l1 := make([]int, 1)
	l1[0] = 1
	l2 := make([]int, 1)
	l2[0] = 1

	if x > 1 {
		t := x
		for t < bound {
			l1 = append(l1, t)
			t *= x
		}
	}

	if y > 1 {
		t := y
		for t < bound {
			l2 = append(l2, t)
			t *= y
		}
	}

	rstMap := make(map[int]bool)
	for _, a := range l1 {
		for _, b := range l2 {
			t := a + b
			if t <= bound {
				rstMap[t] = true
			}
		}
	}

	rst := make([]int, len(rstMap))
	i := 0
	for k := range rstMap {
		rst[i] = k
		i++
	}

	return rst
}
