package mycode

// faster bit count: https://tech.liuchao.me/2016/11/count-bits-of-integer/
func hammingDistance(x int, y int) int {
	t := x ^ y
	t = t - ((t >> 1) & 033333333333) - ((t >> 2) & 011111111111)
	return ((t + (t >> 3)) & 030707070707) % 63
}
