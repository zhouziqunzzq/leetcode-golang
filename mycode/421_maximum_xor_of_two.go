package mycode

// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/discuss/91049/Java-O(n)-solution-using-bit-manipulation-and-HashMap
func findMaximumXOR(nums []int) int {
	maxRst := 0
	mask := 0

	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		set := make(map[int]bool)
		for _, n := range nums {
			set[n&mask] = true
		}

		greedyTry := maxRst | (1 << i)
		for a, _ := range set {
			b := greedyTry ^ a
			if _, ok := set[b]; ok {
				// a XOR b is greedyTry, update maxRst
				maxRst = greedyTry
				break
			}
		}
	}

	return maxRst
}
