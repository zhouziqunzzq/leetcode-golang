package mycode

func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}

	return num == 0
}
