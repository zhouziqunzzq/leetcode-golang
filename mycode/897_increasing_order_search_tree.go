package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	t := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	cur := &t
	increasingBSTHelper(root, &cur)
	return t.Right
}

func increasingBSTHelper(n *TreeNode, cur **TreeNode) {
	if n == nil {
		return
	}

	increasingBSTHelper(n.Left, cur)
	n.Left = nil
	(*cur).Right = n
	*cur = n
	increasingBSTHelper(n.Right, cur)
}
