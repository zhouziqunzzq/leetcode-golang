package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/merge-k-sorted-lists/solution/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	interval := 1

	for interval < n {
		for i := 0; i < n-interval; i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{
		0,
		nil,
	}
	p := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l1
			l1 = p.Next.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return head.Next
}
