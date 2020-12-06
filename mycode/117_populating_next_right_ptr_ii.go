package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/discuss/37828/O(1)-space-O(n)-complexity-Iterative-Solution
func connect(root *Node) *Node {
	var curLevelP, nextLevelP, nextLevelHeadP *Node
	curLevelP = root
	nextLevelHeadP = &Node{}
	nextLevelP = nextLevelHeadP

	for curLevelP != nil {
		if curLevelP.Left != nil {
			nextLevelP.Next = curLevelP.Left
			nextLevelP = nextLevelP.Next
		}
		if curLevelP.Right != nil {
			nextLevelP.Next = curLevelP.Right
			nextLevelP = nextLevelP.Next
		}
		if curLevelP.Next != nil {
			// move to next node in current level
			curLevelP = curLevelP.Next
		} else {
			// move down to next level
			curLevelP = nextLevelHeadP.Next
			nextLevelHeadP.Next = nil
			nextLevelP = nextLevelHeadP
		}
	}

	return root
}
