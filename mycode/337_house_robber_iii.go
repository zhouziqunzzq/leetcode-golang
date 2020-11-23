package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/house-robber-iii/solution/
// Solution 1: recursion
func rob(root *TreeNode) int {
	r, nr := rob3Helper(root)
	return maxInt(r, nr)
}

// return (rob this node, don't rob this node)
func rob3Helper(node *TreeNode) (r, nr int) {
	if node == nil {
		return 0, 0
	}

	lr, lnr := rob3Helper(node.Left)
	rr, rnr := rob3Helper(node.Right)

	r = node.Val + lnr + rnr
	nr = maxInt(lr, lnr) + maxInt(rr, rnr)
	return
}
