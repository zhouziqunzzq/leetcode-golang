package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	rst, _ := subtreeWithAllDeepestHelper(root, 0)
	return rst
}

func subtreeWithAllDeepestHelper(
	node *TreeNode,
	curHeight int,
) (subTreeRoot *TreeNode, maxHeight int) {
	if node == nil {
		subTreeRoot = nil
		maxHeight = curHeight
		return
	}

	lRoot, lHeight := subtreeWithAllDeepestHelper(node.Left, curHeight+1)
	rRoot, rHeight := subtreeWithAllDeepestHelper(node.Right, curHeight+1)

	if lHeight == rHeight {
		subTreeRoot = node
		maxHeight = lHeight
	} else if lHeight > rHeight {
		subTreeRoot = lRoot
		maxHeight = lHeight
	} else {
		subTreeRoot = rRoot
		maxHeight = rHeight
	}
	return
}
