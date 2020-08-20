package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// step 1: find the middle of the list
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// step 2: reverse mid...end
	preMid := p1
	preCur := p1.Next
	for preCur.Next != nil {
		cur := preCur.Next
		preCur.Next = cur.Next
		cur.Next = preMid.Next
		preMid.Next = cur
	}

	// step 3: start reorder
	p1 = head
	p2 = preMid.Next
	for p1 != preMid {
		preMid.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMid.Next
	}
}
