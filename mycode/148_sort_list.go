package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// implement a top-down merge sort
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := sortListGetMidAndCut(head)
	left := sortList(head)
	right := sortList(mid)
	return sortListMerge(left, right)
}

func sortListMerge(l1, l2 *ListNode) *ListNode {
	dummyHead := ListNode{}
	cur := &dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return dummyHead.Next
}

func sortListGetMidAndCut(head *ListNode) *ListNode {
	var midPrev *ListNode = nil
	for head != nil && head.Next != nil {
		if midPrev == nil {
			midPrev = head
		} else {
			midPrev = midPrev.Next
		}
		head = head.Next.Next
	}

	mid := midPrev.Next
	midPrev.Next = nil
	return mid
}
