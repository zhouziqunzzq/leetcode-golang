package mycode

func plusOne(digits []int) []int {
	carry := 0
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			carry = 1
			digits[i] -= 10
		} else {
			carry = 0
			break
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
