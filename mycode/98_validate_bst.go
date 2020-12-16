package mycode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt32, math.MaxInt32)
}

func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val < min || node.Val > max {
		return false
	} else {
		return isValidBSTHelper(node.Left, min, node.Val-1) &&
			isValidBSTHelper(node.Right, node.Val+1, max)
	}
}
