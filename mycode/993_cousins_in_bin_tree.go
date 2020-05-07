package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousinsDFS(node *TreeNode, p *TreeNode, depth uint, target int) (*TreeNode, *TreeNode, uint) {
	if node == nil {
		return nil, p, 0
	}

	if node.Val == target {
		return node, p, depth
	}

	if ln, lp, ld := isCousinsDFS(node.Left, node, depth+1, target); ln != nil {
		return ln, lp, ld
	} else {
		return isCousinsDFS(node.Right, node, depth+1, target)
	}
}

func isCousins(root *TreeNode, x int, y int) bool {
	xn, xp, xd := isCousinsDFS(root, nil, 0, x)
	yn, yp, yd := isCousinsDFS(root, nil, 0, y)

	return xn != nil && yn != nil && xd == yd && xp != nil && yp != nil && xp.Val != yp.Val
}
