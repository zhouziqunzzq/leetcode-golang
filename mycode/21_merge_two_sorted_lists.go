package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	now := sentinel

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			now.Next = l1
			l1 = l1.Next
		} else {
			now.Next = l2
			l2 = l2.Next
		}
		now = now.Next
	}
	if l1 != nil {
		now.Next = l1
	} else if l2 != nil {
		now.Next = l2
	}

	return sentinel.Next
}
