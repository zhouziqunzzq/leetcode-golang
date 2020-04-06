package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(lc *TreeNode, rc *TreeNode) bool {
	if lc == nil && rc == nil {
		return true
	}
	if lc != nil && rc != nil {
		return lc.Val == rc.Val &&
			isMirror(lc.Left, rc.Right) &&
			isMirror(lc.Right, rc.Left)
	}
	return false
}
