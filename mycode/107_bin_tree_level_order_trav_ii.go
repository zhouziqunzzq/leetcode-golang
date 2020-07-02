package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	rst := make([][]int, 0)
	levelOrderBottomHelper(root, 0, &rst)
	reverseSlice(rst)
	return rst
}

// use dfs to achieve level-order traversal
func levelOrderBottomHelper(p *TreeNode, level int, rst *[][]int) {
	if p == nil {
		return
	}
	if len(*rst) < level+1 { // level(root) = 0
		*rst = append(*rst, make([]int, 0))
	}
	levelOrderBottomHelper(p.Left, level+1, rst)
	levelOrderBottomHelper(p.Right, level+1, rst)
	(*rst)[level] = append((*rst)[level], p.Val)
}
