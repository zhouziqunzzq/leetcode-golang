package mycode

// https://leetcode.com/problems/single-number-iii/discuss/68900/Accepted-C%2B%2BJava-O(n)-time-O(1)-space-Easy-Solution-with-Detail-Explanations
func singleNumber3(nums []int) []int {
	// first pass
	t := 0
	for _, n := range nums {
		t ^= n
	}

	// get the rightmost set bit
	t &= -t

	// second pass
	rst := make([]int, 2)
	for _, n := range nums {
		if n&t > 0 {
			rst[0] ^= n
		} else {
			rst[1] ^= n
		}
	}

	return rst
}
