package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbersHelper(n *TreeNode, v int) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return v*10 + n.Val
	}
	return sumNumbersHelper(n.Left, v*10+n.Val) + sumNumbersHelper(n.Right, v*10+n.Val)
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}
