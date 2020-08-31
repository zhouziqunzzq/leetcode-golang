package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// we found a match
		if root.Left == nil { // Case 1: d_out = 1
			return root.Right
		} else if root.Right == nil { // Case 1: d_out = 1
			return root.Left
		} else { // Case 2: d_out = 2
			minNode := root.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}

	return root
}
