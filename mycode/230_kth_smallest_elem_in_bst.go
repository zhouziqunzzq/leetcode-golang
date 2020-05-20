package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// ===========Recursive solution=============

func kthSmallestInOrderTraversal(node *TreeNode, k int, currentK *int) (int, bool) {
	if node.Left != nil {
		if l, hasAns := kthSmallestInOrderTraversal(node.Left, k, currentK); hasAns {
			return l, true
		}
	}

	*currentK++
	if *currentK == k {
		return node.Val, true
	}

	if node.Right != nil {
		return kthSmallestInOrderTraversal(node.Right, k, currentK)
	} else {
		return 0, false
	}
}

func kthSmallest(root *TreeNode, k int) int {
	currentK := 0
	ans, _ := kthSmallestInOrderTraversal(root, k, &currentK)
	return ans
}

// ===========Iterative solution=============

func kthSmallestIterative(root *TreeNode, k int) int {

	stack := make([]*TreeNode, 0)

	for ; root != nil || len(stack) > 0; {
		for ; root != nil; {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		if k == 1 {
			return root.Val
		}
		k--
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return -1
}
