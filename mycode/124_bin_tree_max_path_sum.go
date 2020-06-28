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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func postTraversal(node *TreeNode) (maxInner, maxToRoot int) {
	if node.Left == nil && node.Right == nil {
		maxInner = node.Val
		maxToRoot = node.Val
		return
	}

	maxInnerL, maxToRootL, maxInnerR, maxToRootR :=
	 	math.MinInt32, 0, math.MinInt32, 0
	if node.Left != nil {
		maxInnerL, maxToRootL = postTraversal(node.Left)
	}
	if node.Right != nil {
		maxInnerR, maxToRootR = postTraversal(node.Right)
	}
	maxInner = maxInt3(
		node.Val,
		maxInt(maxInnerL, maxInnerR),
		maxInt(0, maxToRootL) + node.Val + maxInt(0, maxToRootR))
	maxToRoot = maxInt3(
		node.Val,
		maxToRootL + node.Val,
		maxToRootR + node.Val)
	return
}

func maxPathSum(root *TreeNode) int {
	maxInner, maxToRoot := postTraversal(root)
	return maxInt(maxInner, maxToRoot)
}
