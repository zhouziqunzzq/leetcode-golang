package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution 1: using O(n) space with in-order traverse
// https://leetcode.com/problems/recover-binary-search-tree/discuss/32535/No-Fancy-Algorithm-just-Simple-and-Powerful-In-Order-Traversal
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	recoverTreeInOrderTraverse(root, &first, &second, &prev)

	// swap the value of first and second
	first.Val, second.Val = second.Val, first.Val
}

// first: first misplaced elem to be found
// second: second misplaced elem to be found
// prev: the previous node in in-order traverse
func recoverTreeInOrderTraverse(node *TreeNode, first, second, prev **TreeNode) {
	if node == nil {
		return
	}

	recoverTreeInOrderTraverse(node.Left, first, second, prev)

	if *prev != nil && (*prev).Val >= node.Val {
		if *first == nil {
			*first = *prev
		}
		if *first != nil {
			*second = node
		}
	}
	*prev = node

	recoverTreeInOrderTraverse(node.Right, first, second, prev)
}
