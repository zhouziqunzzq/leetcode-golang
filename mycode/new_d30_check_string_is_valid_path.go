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

func isValidSequenceDFS(node *TreeNode, arr []int, pos int) bool {
	if pos >= len(arr) || node.Val != arr[pos] {
		return false
	}

	if node.Left == nil && node.Right == nil {
		if pos == len(arr) - 1 {
			return true
		} else {
			return false
		}
	}

	return (node.Left != nil && isValidSequenceDFS(node.Left, arr, pos+1)) ||
		(node.Right != nil && isValidSequenceDFS(node.Right, arr, pos+1))
}

func isValidSequence(root *TreeNode, arr []int) bool {
	return isValidSequenceDFS(root, arr, 0)
}
