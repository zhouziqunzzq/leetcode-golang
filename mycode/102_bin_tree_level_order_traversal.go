package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	q := make([]*TreeNode, 1)
	q[0] = root

	for len(q) > 0 {
		n := len(q)
		ans = append(ans, make([]int, n))
		for i := 0; i < n; i++ {
			ans[len(ans)-1][i] = q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}

	return ans
}
