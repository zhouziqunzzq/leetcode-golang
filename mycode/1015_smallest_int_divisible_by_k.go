package mycode

// https://leetcode.com/problems/smallest-integer-divisible-by-k/solution/
func smallestRepunitDivByK(K int) int {
	remainder := 0
	for i := 1; i <= K; i++ {
		remainder = (remainder*10 + 1) % K
		if remainder == 0 {
			return i
		}
	}
	return -1
}
