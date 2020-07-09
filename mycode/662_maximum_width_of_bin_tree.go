package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func widthOfBinaryTree(root *TreeNode) int {
	begin := make([]int, 0)
	end := make([]int, 0)
	return widthOfBinaryTreeHelper(root, 0, 1, &begin, &end)
}

// dfs
// idx is the index of the current node if we convert the bin-tree to an array
func widthOfBinaryTreeHelper(p *TreeNode, level, idx int, begin, end *[]int) int {
	if p == nil {
		return 0
	}

	if len(*begin) == level {
		*begin = append(*begin, idx)
		*end = append(*end, idx)
	} else {
		(*end)[level] = idx
	}

	cur := (*end)[level] - (*begin)[level] + 1
	left := widthOfBinaryTreeHelper(p.Left, level+1, 2*idx, begin, end)
	right := widthOfBinaryTreeHelper(p.Right, level+1, 2*idx+1, begin, end)
	return maxInt3(cur, left, right)
}
