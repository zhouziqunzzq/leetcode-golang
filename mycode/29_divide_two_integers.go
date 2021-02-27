package mycode

import "math"

// https://leetcode.com/problems/divide-two-integers/discuss/142849/C%2B%2BJavaPython-Should-Not-Use-%22long%22-Int
// perform division using subtraction and bit manipulation
func divide(dividend int, divisor int) int {
	// handle overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	a := uint32(absInt(dividend))
	b := uint32(absInt(divisor))
	var ans int32 = 0

	for x := 31; x >= 0; x-- {
		if (a >> x) >= b {
			a -= b << x // subtract possible max amount of multiple of b from a
			ans += 1 << x
		}
	}

	if (dividend > 0) == (divisor > 0) {
		return int(ans)
	} else {
		return int(-ans)
	}
}
