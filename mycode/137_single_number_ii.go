package mycode

// https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers
// k = 3, p = 1, k_bin = '11', log(k) = 2
// log(k) = 2, so we need 2 counters: x1, x2
// 2^log(k) = 4 != k, so we need a mask, mask = ~(x1 & x2)
// p' = p % k = 1, p'_bin = '01', so x1 should be the result
func singleNumber(nums []int) int {
	x1, x2, mask := 0, 0, 0

	for _, n := range nums {
		x2 ^= x1 & n
		x1 ^= n
		mask = ^(x1 & x2)
		x2 &= mask
		x1 &= mask
	}

	// we can return x1 or just (x1 | x2)
	return x1
}
