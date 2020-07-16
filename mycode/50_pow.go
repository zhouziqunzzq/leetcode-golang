package mycode

import "math"

// Solution: divide-and-conquer
func myPow(x float64, n int) float64 {
	// base case
	if n == 0 {
		return 1.0
	}

	// deal with min int32
	if n == math.MinInt32 {
		x *= x
		n /= 2
	}

	// now we can safely let n = -n
	if n < 0 {
		n = -n
		x = 1.0 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, n/2)
	}
}
