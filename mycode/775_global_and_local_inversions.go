package mycode

// https://leetcode.com/problems/global-and-local-inversions/discuss/113656/My-3-lines-C%2B%2B-Solution
// Key insight: All local inversions are global, but not the opposite.
func isIdealPermutation(A []int) bool {
	for i, a := range A {
		if a-i > 1 || i-a > 1 {
			// there is a global inversion that is not a local inversion
			return false
		}
	}
	return true
}
