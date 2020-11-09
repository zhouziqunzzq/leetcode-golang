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
func maxAncestorDiff(root *TreeNode) int {
	_, _, ans := maxAncestorDiffHelper(root)
	return ans
}

func maxAncestorDiffHelper(node *TreeNode) (minVal, maxVal, maxDiff int) {
	if node == nil {
		minVal = 1e5 + 10
		maxVal = -1
		maxDiff = 0
		return
	}

	lMinVal, lMaxVal, lMaxDiff := maxAncestorDiffHelper(node.Left)
	rMinVal, rMaxVal, rMaxDiff := maxAncestorDiffHelper(node.Right)

	minVal = minInt(lMinVal, rMinVal)
	maxVal = maxInt(lMaxVal, rMaxVal)
	if minVal == 1e5+10 {
		// current node is a leaf
		minVal = node.Val
		maxVal = node.Val
		maxDiff = 0
		return
	} else {
		maxDiff = maxInt(lMaxDiff, rMaxDiff)
		maxDiff = maxInt(maxDiff, int(math.Abs(float64(node.Val-minVal))))
		maxDiff = maxInt(maxDiff, int(math.Abs(float64(maxVal-node.Val))))
		minVal = minInt(minVal, node.Val)
		maxVal = maxInt(maxVal, node.Val)
		return
	}
}
