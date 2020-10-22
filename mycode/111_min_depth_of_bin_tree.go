package mycode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return minDepthHelper(root, 1, math.MaxInt32)
}

func minDepthHelper(node *TreeNode, curDepth, curMin int) int {
	// leaf
	if node.Left == nil && node.Right == nil {
		return curDepth
	}

	// pruning
	if curDepth >= curMin {
		return curDepth
	}

	if node.Left != nil {
		l := minDepthHelper(node.Left, curDepth+1, curMin)
		curMin = minInt(curMin, l)
	}
	if node.Right != nil {
		r := minDepthHelper(node.Right, curDepth+1, curMin)
		curMin = minInt(curMin, r)
	}
	return curMin
}
