package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	sentinel := &ListNode{
		Val:  0,
		Next: head,
	}

	pp, cp, np := sentinel, head, head.Next
	removeCurr := false

	for cp != nil && np != nil {
		// reset flag
		removeCurr = false

		for np != nil && cp.Val == np.Val {
			removeCurr = true
			np = np.Next
		}

		if removeCurr {
			pp.Next = np
			cp = np
			if np == nil {
				break
			} else {
				np = np.Next
			}
		} else {
			pp = cp
			cp = np
			np = np.Next
		}
	}

	return sentinel.Next
}
