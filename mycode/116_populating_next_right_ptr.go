package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/discuss/37472/A-simple-accepted-solution
// Solution 1: iterative method with 2 pointers
func connect116(root *Node) *Node {
	if root == nil {
		return root
	}

	pre, cur := root, root
	for pre.Left != nil {
		cur = pre

		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}

		pre = pre.Left
	}

	return root
}

// Solution 2: recursive method
func connectRecursive(root *Node) *Node {
	if root == nil {
		return root
	}

	connectRecursiveHelper(root.Left, root.Right)
	return root
}

func connectRecursiveHelper(n1, n2 *Node) {
	if n1 == nil {
		return
	}

	n1.Next = n2
	connectRecursiveHelper(n1.Right, n2.Left)
	connectRecursiveHelper(n1.Left, n1.Right)
	connectRecursiveHelper(n2.Left, n2.Right)
}
