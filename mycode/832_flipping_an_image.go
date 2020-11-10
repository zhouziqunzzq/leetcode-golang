package mycode

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		for i := 0; i < (len(a)+1)/2; i++ {
			// flip
			a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
			// inverse
			if a[i] == 0 {
				a[i] = 1
			} else {
				a[i] = 0
			}
			if i == len(a)-1-i {
				continue
			}
			if a[len(a)-1-i] == 0 {
				a[len(a)-1-i] = 1
			} else {
				a[len(a)-1-i] = 0
			}
		}
	}

	return A
}
