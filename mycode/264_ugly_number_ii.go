package mycode

func nthUglyNumber(n int) int {
	// base cases
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	uglyNums := make([]int, n)
	uglyNums[0] = 1
	p1, p2, p3 := 0, 0, 0

	for i := 1; i < n; i++ {
		uglyNums[i] = minInt3(uglyNums[p1]*2, uglyNums[p2]*3, uglyNums[p3]*5)
		if uglyNums[i] == uglyNums[p1]*2 {
			p1++
		}
		if uglyNums[i] == uglyNums[p2]*3 {
			p2++
		}
		if uglyNums[i] == uglyNums[p3]*5 {
			p3++
		}
	}

	return uglyNums[n-1]
}
