package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumRootToLeafHelper(root, 0)
}

func sumRootToLeafHelper(node *TreeNode, curSum int) int {
	if node.Left == nil && node.Right == nil {
		return (curSum << 1) + node.Val
	}

	sum := 0
	if node.Left != nil {
		sum += sumRootToLeafHelper(node.Left, (curSum<<1)+node.Val)
	}
	if node.Right != nil {
		sum += sumRootToLeafHelper(node.Right, (curSum<<1)+node.Val)
	}
	return sum
}
