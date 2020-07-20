package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// special case for head
	for head != nil && head.Val == val {
		head = head.Next
	}

	pPrev, p := head, head
	for p != nil {
		if p.Val == val {
			pPrev.Next = p.Next
		} else {
			pPrev = p
		}
		p = p.Next
	}

	return head
}
