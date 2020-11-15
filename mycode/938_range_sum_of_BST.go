package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	// base case
	if root == nil {
		return 0
	}

	if low == root.Val {
		return root.Val + rangeSumBST(root.Right, root.Val, high)
	} else if high == root.Val {
		return root.Val + rangeSumBST(root.Left, low, root.Val)
	} else if low < root.Val && root.Val < high {
		return root.Val +
			rangeSumBST(root.Left, low, root.Val) +
			rangeSumBST(root.Right, root.Val, high)
	} else if high < root.Val {
		return rangeSumBST(root.Left, low, high)
	} else { // root.Val < low
		return rangeSumBST(root.Right, low, high)
	}
}
