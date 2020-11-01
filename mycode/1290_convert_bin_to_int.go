package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	ans := 0
	p := head

	for p != nil {
		ans <<= 1
		if p.Val == 1 {
			ans |= 1
		}
		p = p.Next
	}

	return ans
}
