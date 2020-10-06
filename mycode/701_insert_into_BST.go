package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	par, cur := root, root
	isLeft := false
	for cur != nil {
		par = cur
		if val < cur.Val {
			isLeft = true
			cur = cur.Left
		} else {
			isLeft = false
			cur = cur.Right
		}
	}

	if isLeft {
		par.Left = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		par.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	return root
}
