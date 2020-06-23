package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodesHelperHeight(n *TreeNode) int {
	if n == nil {
		return -1
	} else {
		return 1 + countNodesHelperHeight(n.Left)
	}
}

func countNodes(root *TreeNode) int {
	cnt := 0
	h := countNodesHelperHeight(root)
	hRight := 0

	for root != nil {
		hRight = countNodesHelperHeight(root.Right)
		if hRight == h-1 {
			// left subtree is a full tree of h-1
			cnt += 1 << h
			root = root.Right
		} else {
			// left subtree is not full, but right subtree is a full tree of h-2
			cnt += 1 << (h - 1)
			root = root.Left
		}
		// go down a level
		h--
	}

	return cnt
}
