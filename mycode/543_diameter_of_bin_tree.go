package mycode

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

//func maxInt(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

// return: (maxDepth, maxDiameter)
func diameterOfBinaryTreeHelper(node *TreeNode) (int, int) {
	// base case
	if node.Left == nil && node.Right == nil {
		return 0, 0
	}

	if node.Left != nil && node.Right == nil {
		lDepth, lDiameter := diameterOfBinaryTreeHelper(node.Left)
		return lDepth+1, maxInt(lDiameter, lDepth+1)
	} else if node.Left == nil && node.Right != nil {
		rDepth, rDiameter := diameterOfBinaryTreeHelper(node.Right)
		return rDepth+1, maxInt(rDiameter, rDepth+1)
	} else {
		lDepth, lDiameter := diameterOfBinaryTreeHelper(node.Left)
		rDepth, rDiameter := diameterOfBinaryTreeHelper(node.Right)
		return maxInt(lDepth, rDepth)+1, maxInt(maxInt(lDiameter, rDiameter), lDepth+rDepth+2)
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		_, ans := diameterOfBinaryTreeHelper(root)
		return ans
	}
}
