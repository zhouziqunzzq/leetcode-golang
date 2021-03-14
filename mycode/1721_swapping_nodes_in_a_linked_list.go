package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/discuss/1009800/C%2B%2BJP3-One-Pass
// Brilliant 1-pass solution using 2 ptrs!
func swapNodes(head *ListNode, k int) *ListNode {
	var p1, p2 *ListNode
	p := head

	for p != nil {
		if p2 != nil {
			p2 = p2.Next
		}
		k--
		if k == 0 {
			p1 = p
			p2 = head
		}
		p = p.Next
	}

	//goland:noinspection ALL
	p1.Val, p2.Val = p2.Val, p1.Val

	return head
}
