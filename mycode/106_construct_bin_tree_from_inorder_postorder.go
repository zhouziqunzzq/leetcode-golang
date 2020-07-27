package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	pInorder := len(inorder) - 1
	pPostorder := len(postorder) - 1

	return buildTreeHelper(inorder, postorder, nil, &pInorder, &pPostorder)
}

func buildTreeHelper(inorder []int, postorder []int, end *TreeNode, pInorder, pPostorder *int) *TreeNode {
	if *pPostorder < 0 {
		return nil
	}

	// root
	n := &TreeNode{postorder[*pPostorder], nil, nil}
	*pPostorder--

	// right subtree
	if inorder[*pInorder] != n.Val {
		n.Right = buildTreeHelper(inorder, postorder, n, pInorder, pPostorder)
	}

	*pInorder--

	// left subtree
	if end == nil || inorder[*pInorder] != end.Val {
		n.Left = buildTreeHelper(inorder, postorder, end, pInorder, pPostorder)
	}

	return n
}
