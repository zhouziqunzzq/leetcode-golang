package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/49785/Java-solution-without-knowing-the-difference-in-len!
// This solution is SO smart and concise! love it
// By swapping two ptrs at the end of each loop, it's guaranteed to finish in at most 2 loops.
// Proof:
// Suppose the length of list A is a + c and that of list B is b + c.
// (where c is the length of the intersected part, and let's assume a != b for now)
// Total steps of ptr_a traveled: a + c + b + c
// Total steps of ptr_b traveled: b + c + a + c
// Note that the first 3 terms in above two expressions, (a + c + b) = (b + c + a).
// So, ptr_a and ptr_b will meet at the intersection in the 2nd loop!
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB // reach the end of A, swap ptr a to B
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA // reach the end of B, swap ptr b to A
		} else {
			b = b.Next
		}
	}

	return a
}
