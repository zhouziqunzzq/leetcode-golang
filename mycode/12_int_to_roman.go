package mycode

func intToRoman(num int) string {
	ans := ""
	n := num / 1000
	if n > 0 {
		for i := 0; i < n; i++ {
			ans += "M"
		}
	}
	num -= n * 1000

	n = num / 100
	if n > 0 {
		if n <= 3 {
			for i := 0; i < n; i++ {
				ans += "C"
			}
		} else if n == 4 {
			ans += "CD"
		} else if n <= 8 {
			ans += "D"
			for i := 0; i < n-5; i++ {
				ans += "C"
			}
		} else {
			ans += "CM"
		}
	}
	num -= n * 100

	n = num / 10
	if n > 0 {
		if n <= 3 {
			for i := 0; i < n; i++ {
				ans += "X"
			}
		} else if n == 4 {
			ans += "XL"
		} else if n <= 8 {
			ans += "L"
			for i := 0; i < n-5; i++ {
				ans += "X"
			}
		} else {
			ans += "XC"
		}
	}
	num -= n * 10

	n = num % 10
	if n > 0 {
		if n <= 3 {
			for i := 0; i < n; i++ {
				ans += "I"
			}
		} else if n == 4 {
			ans += "IV"
		} else if n <= 8 {
			ans += "V"
			for i := 0; i < n-5; i++ {
				ans += "I"
			}
		} else {
			ans += "IX"
		}
	}

	return ans
}
