package mycode

// https://leetcode.com/problems/beautiful-arrangement-ii/solution/
func constructArray(n int, k int) []int {
	ans := make([]int, n)
	p := 0

	// first part: every two adj elem have the same diff of 1 (1, 2, 3, ..., n)
	for i := 1; i <= n-k-1; i++ {
		ans[p] = i
		p++
	}

	// second part: every two adj elem have a different diff (1, n, 2, n-1, 3, n-2, ...)
	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			ans[p] = n - k + i/2
		} else {
			ans[p] = n - i/2
		}
		p++
	}

	return ans
}
