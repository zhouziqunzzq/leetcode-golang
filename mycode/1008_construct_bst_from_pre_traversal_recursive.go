package mycode

import "math"

// Another solution with recursive method. Extremely concise. Awesome! See:
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252232/JavaC%2B%2BPython-O(N)-Solution

func bstFromPreorderHelper(preorder []int, i *int, upperBound int) *TreeNode {
	if *i == len(preorder) || preorder[*i] > upperBound {
		return nil
	}
	root := &TreeNode{preorder[*i], nil, nil}
	*i++
	root.Left = bstFromPreorderHelper(preorder, i, root.Val)
	root.Right = bstFromPreorderHelper(preorder, i, upperBound)

	return root
}

func bstFromPreorderRecursive(preorder []int) *TreeNode {
	i := 0
	return bstFromPreorderHelper(preorder, &i, math.MaxInt32)
}
