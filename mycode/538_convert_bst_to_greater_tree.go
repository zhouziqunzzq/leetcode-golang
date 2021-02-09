package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	convertBSTHelper(root, 0)
	return root
}

func convertBSTHelper(node *TreeNode, sumFromRight int) int {
	if node == nil {
		return 0
	}

	sr := convertBSTHelper(node.Right, sumFromRight)
	sl := convertBSTHelper(node.Left, node.Val+sr+sumFromRight)
	s := sr + sl + node.Val
	node.Val += sumFromRight + sr
	return s
}

// https://leetcode.com/problems/convert-bst-to-greater-tree/solution/
// We can also use Morris traversal to achieve constant space.
