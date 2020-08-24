package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesHelper(root, false)
}

func sumOfLeftLeavesHelper(node *TreeNode, isLeft bool) int {
	// base case 1
	if node == nil {
		return 0
	}

	// base case 2
	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}

	return sumOfLeftLeavesHelper(node.Left, true) +
		sumOfLeftLeavesHelper(node.Right, false)
}
