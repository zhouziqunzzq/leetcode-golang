package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/binary-tree-cameras/
// Solution 1: DP
// Solution 2: Greedy
func minCameraCover(root *TreeNode) int {
	_, dp1, dp2 := minCameraCoverDpHelper(root)
	return minInt(dp1, dp2)
}

func minCameraCoverDpHelper(node *TreeNode) (dp0, dp1, dp2 int) {
	// 0: Strict subtree; All nodes below this are covered, but not this one.
	// 1: Normal subtree; All nodes below and incl this are covered, no camera here.
	// 2: Placed camera;  All nodes below this are covered, placed camera here.
	if node == nil {
		return 0, 0, 1e6
	}

	l0, l1, l2 := minCameraCoverDpHelper(node.Left)
	r0, r1, r2 := minCameraCoverDpHelper(node.Right)

	dp0 = l1 + r1
	dp1 = minInt(l2+minInt(r1, r2), r2+minInt(l1, l2))
	dp2 = 1 + minInt3(l0, l1, l2) + minInt3(r0, r1, r2)

	return
}

func minCameraCoverGreedy(root *TreeNode) int {
	ans := 0
	covered := make(map[*TreeNode]bool)
	covered[nil] = true
	minCameraCoverGreedyHelper(root, nil, &ans, covered)
	return ans
}

func minCameraCoverGreedyHelper(node, parent *TreeNode, ans *int, covered map[*TreeNode]bool) {
	if node != nil {
		minCameraCoverGreedyHelper(node.Left, node, ans, covered)
		minCameraCoverGreedyHelper(node.Right, node, ans, covered)

		_, nCovered := covered[node]
		_, lCovered := covered[node.Left]
		_, rCovered := covered[node.Right]
		if (parent == nil && !nCovered) || (!lCovered || !rCovered) {
			*ans++
			covered[node] = true
			covered[node.Left] = true
			covered[node.Right] = true
			covered[parent] = true
		}
	}
}
