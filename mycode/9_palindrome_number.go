package mycode

func isPalindrome9(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		digits := make([]int8, 0)
		for x > 0 {
			digits = append(digits, int8(x%10))
			x = x / 10
		}

		i, j := 0, len(digits)-1
		for i <= j {
			if digits[i] != digits[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
}
