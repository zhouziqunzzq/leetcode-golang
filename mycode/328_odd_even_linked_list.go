package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	} else {
		// at least 2 nodes
		oddHead, evenHead := head, head.Next
		oddP, evenP := oddHead, evenHead
		oldP := evenHead.Next
		isOdd := true

		for oldP != nil {
			nextP := oldP.Next
			if isOdd {
				oddP.Next = oldP
				oddP = oddP.Next
			} else {
				evenP.Next = oldP
				evenP = evenP.Next
			}
			oldP = nextP
			isOdd = !isOdd
		}

		oddP.Next = evenHead
		evenP.Next = nil
		return oddHead
	}
}
