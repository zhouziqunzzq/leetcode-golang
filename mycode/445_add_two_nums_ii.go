package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	sum := 0
	rst := &ListNode{
		Val:  0,
		Next: nil,
	}
	for len(s1) > 0 || len(s2) > 0 {
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		rst.Val = sum % 10
		nextRst := &ListNode{
			Val:  sum / 10,
			Next: rst,
		}
		rst = nextRst
		sum /= 10
	}

	if rst.Val == 0 {
		return rst.Next
	} else {
		return rst
	}
}
