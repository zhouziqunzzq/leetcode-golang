package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	pp := node
	pn := node.Next

	for pn.Next != nil {
		pp.Val = pn.Val
		pp = pn
		pn = pn.Next
	}

	pp.Val = pn.Val
	pp.Next = nil
}
