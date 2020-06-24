package mycode

// math is here:
// https://leetcode.com/problems/unique-binary-search-trees/discuss/31666/DP-Solution-in-6-lines-with-explanation.-F(i-n)-G(i-1)-*-G(n-i)
// https://www.geeksforgeeks.org/program-nth-catalan-number/
func numTrees(n int) int {
	opt := make([]int, n+1)

	// base case
	opt[0], opt[1] = 1, 1

	// fill the table
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			opt[i] += opt[j-1] * opt[i-j]
		}
	}

	return opt[n]
}
