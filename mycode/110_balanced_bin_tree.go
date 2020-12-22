package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, ans := isBalancedHelper(root, 0)
	return ans
}

// get the height of a bin-tree rooted at node
func isBalancedHelper(node *TreeNode, curHeight int) (h int, isBalanced bool) {
	if node == nil {
		h = curHeight
		isBalanced = true
		return
	}

	lh, lb := isBalancedHelper(node.Left, curHeight+1)
	rh, rb := isBalancedHelper(node.Right, curHeight+1)
	h = maxInt(lh, rh)
	diff := lh - rh
	isBalanced = lb && rb && -1 <= diff && diff <= 1
	return
}
