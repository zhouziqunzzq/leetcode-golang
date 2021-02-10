package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	if head == nil {
		return nil
	}

	var cur, next *Node138

	// copy nodes and interweave copied nodes with original nodes
	cur = head
	for cur != nil {
		next = cur.Next
		newNode := &Node138{Val: cur.Val, Next: next}
		cur.Next = newNode
		cur = next
	}

	// set Random pointer for copied nodes
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// restore original list
	newHead := head.Next
	cur = head
	for cur != nil {
		next = cur.Next.Next
		if next != nil {
			cur.Next.Next = next.Next
		}
		cur.Next = next
		cur = next
	}

	return newHead
}
