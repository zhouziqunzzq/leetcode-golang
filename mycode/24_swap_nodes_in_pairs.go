package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, post, p1, p2, newHead *ListNode
	p1 = head
	p2 = p1.Next
	post = p2.Next

	for {
		// adjust nodes
		if pre == nil {
			// first pair of nodes
			newHead = p2
		} else {
			pre.Next = p2
		}
		p2.Next = p1
		p1.Next = post

		// check terminating condition
		if post == nil || post.Next == nil {
			break
		}

		// forward pointers
		pre = p1
		p1 = pre.Next
		p2 = p1.Next
		post = p2.Next
	}

	return newHead
}
