package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS
func deepestLeavesSum(root *TreeNode) int {
	lastSum := 0
	q := make([]*TreeNode, 1)
	q[0] = root

	for len(q) > 0 {
		lastSum = 0
		kn := len(q)
		for i := 0; i < kn; i++ {
			lastSum += q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[kn:]
	}

	return lastSum
}
