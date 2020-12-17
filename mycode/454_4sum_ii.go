package mycode

// similar to 2sum, we use hashmap to store a+b and iterate c+d to count.
func fourSumCount(A []int, B []int, C []int, D []int) int {
	ab := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			if _, ok := ab[a+b]; !ok {
				ab[a+b] = 1
			} else {
				ab[a+b]++
			}
		}
	}

	rst := 0
	for _, c := range C {
		for _, d := range D {
			if cnt, ok := ab[-(c + d)]; ok {
				rst += cnt
			}
		}
	}

	return rst
}
