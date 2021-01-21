package mycode

// https://leetcode.com/problems/find-the-most-competitive-subsequence/discuss/952786/JavaC%2B%2BPython-One-Pass-Stack-Solution
// using monotonically increasing stack
func mostCompetitive(nums []int, k int) []int {
	s := make([]int, 0, k)
	n := len(nums)

	for i, v := range nums {
		// pop while s.top > v, but make sure there's enough elem remaining in nums!
		for len(s) > 0 && s[len(s)-1] > v && len(s)-1+n-i >= k {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}

	return s
}
