package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	h1 := &ListNode{}
	h2 := &ListNode{}
	p1, p2 := h1, h2
	p := head

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		p = p.Next
	}

	p1.Next = h2.Next
	p2.Next = nil
	return h1.Next
}
