package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/discuss/35476/Share-my-JAVA-solution-1ms-very-short-and-concise.
// Two-ptrs + DFS
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else {
		return sortedListToBSTHelper(head, nil)
	}
}

func sortedListToBSTHelper(head, tail *ListNode) *TreeNode {
	if head == tail { // base case
		return nil
	}

	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// now slow points to mid
	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBSTHelper(head, slow)
	root.Right = sortedListToBSTHelper(slow.Next, tail)
	return root
}
