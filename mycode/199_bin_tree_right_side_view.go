package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	rst := make([]int, 0)
	if root == nil {
		return rst
	}

	q := make([]*TreeNode, 1)
	q[0] = root

	for len(q) > 0 {
		prevLen := len(q)
		rst = append(rst, q[prevLen-1].Val)

		for i := 0; i < prevLen; i++ {
			pn := q[0]
			q = q[1:]
			if pn.Left != nil {
				q = append(q, pn.Left)
			}
			if pn.Right != nil {
				q = append(q, pn.Right)
			}
		}
	}

	return rst
}
