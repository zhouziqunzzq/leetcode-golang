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
func findTilt(root *TreeNode) int {
	_, ans := findTiltHelper(root)
	return ans
}

func findTiltHelper(node *TreeNode) (sum, sumTilt int) {
	if node == nil {
		sum = 0
		sumTilt = 0
		return
	}

	lSum, lTilt := findTiltHelper(node.Left)
	rSum, rTilt := findTiltHelper(node.Right)

	sum = lSum + rSum + node.Val
	sumTilt = lTilt + rTilt + int(math.Abs(float64(lSum-rSum)))
	return
}
