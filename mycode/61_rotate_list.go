package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l := 1
	tail := head
	for tail.Next != nil {
		l++
		tail = tail.Next
	}

	tail.Next = head
	k %= l
	for i := 0; i < l-k; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil

	return head
}
