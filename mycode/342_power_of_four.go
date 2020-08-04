package mycode

// Another solution: https://leetcode.com/problems/power-of-four/discuss/80460/1-line-C%2B%2B-solution-without-confusing-bit-manipulations
func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}
