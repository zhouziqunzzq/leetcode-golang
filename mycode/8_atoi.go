package mycode

import "unicode"

func myAtoi(str string) int {
	const IntMax = int32(^uint32(0) >> 1)
	const IntMin = ^IntMax

	var ans int32 = 0
	isNeg := false
	isParsing := false

	strRune := []rune(str)
	for _, r := range strRune {
		if r == ' ' {
			if isParsing {
				break
			} else {
				continue
			}
		} else if r == '-' && !isParsing {
			isNeg = true
			isParsing = true
		} else if r == '+' && !isParsing {
			isNeg = false
			isParsing = true
		} else if unicode.IsDigit(r) {
			isParsing = true
			i := int32(r - '0')

			if ans >= 0 {
				if (ans > IntMax/10) || (ans == IntMax/10 && i > 7) {
					ans = IntMax
					break
				}
				ans = ans*10 + i
			} else if ans < 0 {
				if (ans < IntMin/10) || (ans == IntMin/10 && i > 8) {
					ans = IntMin
					break
				}
				ans = ans*10 - i
			}
			if isNeg && ans > 0 {
				ans = -ans
			}
		} else {
			break
		}
	}
	return int(ans)
}
