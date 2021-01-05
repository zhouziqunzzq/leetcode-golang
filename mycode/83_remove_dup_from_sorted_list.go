package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesI(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else {
		lastElem := head.Val
		lastP := head
		nowP := head.Next
		for nowP != nil {
			for nowP != nil && nowP.Val == lastElem {
				nowP = nowP.Next
			}
			lastP.Next = nowP
			lastP = nowP
			if nowP != nil {
				lastElem = nowP.Val
			}
		}
	}
	return head
}
