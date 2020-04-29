package mycode

import "math"

func getSum(root *TreeNode, globalMax *int) int {
	if root == nil {
		return 0
	}
	leftSum := maxInt(getSum(root.Left, globalMax), 0)
	rightSum := maxInt(getSum(root.Right, globalMax),0)

	sum := root.Val + leftSum + rightSum

	*globalMax = maxInt(sum, *globalMax)

	return root.Val + maxInt(leftSum, rightSum)
}

func maxPathSum2(root *TreeNode) int {
	globalMax := math.MinInt32
	getSum(root, &globalMax)
	return globalMax
}
