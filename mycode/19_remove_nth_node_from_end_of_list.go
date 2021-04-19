package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/solution/
// two pointers
// the gap between first and second is n nodes apart
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// use dummy here to handle corner case of deleting the 1st node of the original linked list
	dummy := &ListNode{0, head}
	p1, p2 := dummy, dummy

	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}
