package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/solution/
// DFS, try to flip on the fly
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	rst := make([]int, 0)
	idx := 0
	flipMatchVoyageDFS(root, voyage, &rst, &idx)
	if len(rst) > 0 && rst[0] == -1 {
		rst = make([]int, 1)
		rst[0] = -1
	}
	return rst
}

func flipMatchVoyageDFS(node *TreeNode, voyage []int, rst *[]int, idx *int) {
	if node == nil {
		return
	}

	// check root val
	if node.Val != voyage[*idx] {
		// can't get desired pre-order traversal, return [-1]
		*rst = make([]int, 1)
		(*rst)[0] = -1
		*idx++
		return
	}
	*idx++

	// call DFS on subtrees and try to flip if needed
	if *idx < len(voyage) && node.Left != nil && node.Left.Val != voyage[*idx] {
		*rst = append(*rst, node.Val)
		flipMatchVoyageDFS(node.Right, voyage, rst, idx)
		flipMatchVoyageDFS(node.Left, voyage, rst, idx)
	} else {
		flipMatchVoyageDFS(node.Left, voyage, rst, idx)
		flipMatchVoyageDFS(node.Right, voyage, rst, idx)
	}
}
