package mycode

func brokenCalc(X int, Y int) int {
	ans := 0

	for Y != X {
		if Y > X {
			if Y%2 == 0 {
				Y /= 2
				ans++
			} else {
				Y = (Y + 1) / 2
				ans += 2
			}
		} else {
			ans += X - Y
			break
		}
	}

	return ans
}
