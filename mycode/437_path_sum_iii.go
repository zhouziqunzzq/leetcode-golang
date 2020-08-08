package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	m := make(map[int]int)
	m[0] = 1
	return pathSumHelper(root, 0, sum, m)
}

// m: prefixSum -> count
func pathSumHelper(n *TreeNode, curSum, target int, m map[int]int) int {
	if n == nil {
		return 0
	}

	curSum += n.Val

	rst := 0
	if cnt, ok := m[curSum-target]; ok {
		rst += cnt
	}

	m[curSum]++
	rst += (pathSumHelper(n.Left, curSum, target, m) + pathSumHelper(n.Right, curSum, target, m))
	// backtrace
	m[curSum]--

	return rst
}
