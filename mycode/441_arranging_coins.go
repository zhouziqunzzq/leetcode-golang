package mycode

import "math"

func arrangeCoins(n int) int {
	return int(math.Floor((-1 + math.Sqrt(1.0+8.0*float64(n))) / 2.0))
}
