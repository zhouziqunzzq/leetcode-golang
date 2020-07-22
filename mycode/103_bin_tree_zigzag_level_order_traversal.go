package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	rst := make([][]int, 0)
	zigzagLevelOrderHelper(root, 0, &rst)
	for i, r := range rst {
		if i&1 != 0 {
			reverseSlice(r)
		}
	}
	return rst
}

// use dfs to achieve level-order traversal
func zigzagLevelOrderHelper(p *TreeNode, level int, rst *[][]int) {
	if p == nil {
		return
	}
	if len(*rst) < level+1 { // level(root) = 0
		*rst = append(*rst, make([]int, 0))
	}
	zigzagLevelOrderHelper(p.Left, level+1, rst)
	zigzagLevelOrderHelper(p.Right, level+1, rst)
	(*rst)[level] = append((*rst)[level], p.Val)
}
