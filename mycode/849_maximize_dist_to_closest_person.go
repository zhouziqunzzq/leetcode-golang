package mycode

func maxDistToClosest(seats []int) int {
	ans := 0
	l, r := 0, len(seats)-1

	// find the first 1 starting from left
	for seats[l] != 1 {
		l++
	}
	ans = maxInt(ans, l)

	// find the last 1 starting from right
	for seats[r] != 1 {
		r--
	}
	ans = maxInt(ans, len(seats)-r-1)

	// find intervals like 1000...0001
	for i := l; i <= r; i++ {
		cur := 1
		for i <= r && seats[i] == 1 {
			i++
		}
		for i <= r && seats[i] == 0 {
			cur++
			i++
		}
		ans = maxInt(ans, cur/2)
	}

	return ans
}
