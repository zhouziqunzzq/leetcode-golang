package mycode

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  math.MinInt32,
		Next: nil,
	}
	p := head

	for p != nil {
		prev := dummyHead
		cur := dummyHead

		for cur != nil && cur.Val <= p.Val {
			prev = cur
			cur = cur.Next
		}

		nextP := p.Next
		prev.Next = p
		p.Next = cur

		p = nextP
	}

	dummyHead = dummyHead.Next
	return dummyHead
}
