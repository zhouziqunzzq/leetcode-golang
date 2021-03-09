package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := &TreeNode{v, root, nil}
		return newRoot
	}

	addOneRowDFS(root, v, d, 1)
	return root
}

func addOneRowDFS(node *TreeNode, v, d, curD int) {
	if node == nil {
		return
	}

	if curD == d-1 {
		oldL, oldR := node.Left, node.Right
		node.Left = &TreeNode{v, oldL, nil}
		node.Right = &TreeNode{v, nil, oldR}
		return
	}

	addOneRowDFS(node.Left, v, d, curD+1)
	addOneRowDFS(node.Right, v, d, curD+1)
}
