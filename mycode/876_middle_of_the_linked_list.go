package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func middleNode(head *ListNode) *ListNode {
	p := head
	cnt := 0

	for p != nil {
		p = p.Next
		cnt++
	}

	cnt /= 2
	p = head
	for cnt > 0 {
		p = p.Next
		cnt--
	}

	return p
}

// Note: another solution using fast and slow pointers
func middleNodeSlowFast(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
