package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type Node430 struct {
	Val   int
	Prev  *Node430
	Next  *Node430
	Child *Node430
}

func flatten(root *Node430) *Node430 {
	if root != nil {
		flattenSolveHelper(root)
	}
	return root
}

func flattenSolveHelper(p *Node430) *Node430 {
	for {
		if p.Child != nil {
			childLast := flattenSolveHelper(p.Child)
			flattenConnectHelper(childLast, p.Next)
			flattenConnectHelper(p, p.Child)
			p.Child = nil
			if childLast.Next != nil {
				p = childLast.Next
			} else {
				p = childLast
				break
			}
		} else {
			if p.Next != nil {
				p = p.Next
			} else {
				break
			}
		}
	}

	return p
}

func flattenConnectHelper(n1, n2 *Node430) {
	if n1 != nil {
		n1.Next = n2
	}
	if n2 != nil {
		n2.Prev = n1
	}
}
