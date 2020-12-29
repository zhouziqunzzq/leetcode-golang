package mycode

import "math/bits"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	return pseudoPalindromicPathsHelper(root, 0)
}

func pseudoPalindromicPathsHelper(node *TreeNode, prevFlags uint32) int {
	if node == nil {
		return 0
	}

	f := uint32(1 << node.Val)
	curFlags := prevFlags ^ f

	if node.Left == nil && node.Right == nil {
		if bits.OnesCount32(curFlags) > 1 {
			// more than one digits of odd parity count
			return 0
		} else {
			return 1
		}
	}

	return pseudoPalindromicPathsHelper(node.Left, curFlags) +
		pseudoPalindromicPathsHelper(node.Right, curFlags)
}
