package mycode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	p := head
	carry := 0
	for {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
		} else {
			b = 0
		}
		// calc val
		p.Val = a + b + carry
		// handle carry
		if p.Val >= 10 {
			carry = 1
			p.Val -= 10
		} else {
			carry = 0
		}
		// move p
		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) {
			p.Next = &ListNode{0, nil}
			p = p.Next
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		} else {
			if carry != 0 {
				p.Next = &ListNode{1, nil}
			}
			break
		}
	}
	return head
}
