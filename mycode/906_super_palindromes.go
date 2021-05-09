package mycode

import "strconv"
import "strings"

// https://leetcode.com/problems/super-palindromes/solution/
// Math + Brute force

func superpalindromesInRange(left string, right string) int {
	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)
	const MaxHalfVal int = 1e5
	ans := 0

	// count palindromes with odd length
	for k := 1; k < MaxHalfVal; k++ {
		kStr := strconv.Itoa(k)

		// k | k'
		// e.g. k = 123, sb = 12321
		sb := strings.Builder{}
		sb.WriteString(kStr)
		for i := len(kStr) - 2; i >= 0; i-- {
			sb.WriteRune(rune(kStr[i]))
		}

		v, _ := strconv.Atoi(sb.String())
		v *= v
		if v > r {
			break
		}
		if v >= l && isPalindromeInt(v) {
			ans++
		}
	}

	// count palindromes with even length
	for k := 1; k < MaxHalfVal; k++ {
		kStr := strconv.Itoa(k)

		// k | k'
		// e.g. k = 123, sb = 123321
		sb := strings.Builder{}
		sb.WriteString(kStr)
		for i := len(kStr) - 1; i >= 0; i-- {
			sb.WriteRune(rune(kStr[i]))
		}

		v, _ := strconv.Atoi(sb.String())
		v *= v
		if v > r {
			break
		}
		if v >= l && isPalindromeInt(v) {
			ans++
		}
	}

	return ans
}
