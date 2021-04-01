package mycode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/palindrome-linked-list/discuss/64500/11-lines-12-with-restore-O(n)-time-O(1)-space
// Reverse first half and compare with second half.
func isPalindrome(head *ListNode) bool {
	var rev *ListNode
	slow, fast := head, head

	// reverse first half
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev := rev
		next := slow.Next
		rev = slow
		rev.Next = prev
		slow = next
	}

	revTail := slow
	// handle odd number of nodes
	if fast != nil {
		slow = slow.Next
	}

	// compare and restore
	ans := true
	for rev != nil {
		ans = rev.Val == slow.Val && ans
		slow = slow.Next
		prev := rev
		next := rev.Next
		rev.Next = revTail
		revTail = prev
		rev = next
	}

	return ans
}
