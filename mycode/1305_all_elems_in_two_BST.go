package mycode

// https://leetcode.com/problems/all-elements-in-two-binary-search-trees/discuss/464073/C%2B%2B-One-Pass-Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	rst := make([]int, 0)
	s1 := make([]*TreeNode, 0)
	s2 := make([]*TreeNode, 0)
	getAllElementsGoLeft(root1, &s1)
	getAllElementsGoLeft(root2, &s2)

	for len(s1) > 0 || len(s2) > 0 {
		s := &s1
		if len(s1) > 0 && len(s2) > 0 {
			if s1[len(s1)-1].Val > s2[len(s2)-1].Val {
				s = &s2
			}
		} else if len(s2) > 0 {
			s = &s2
		}
		n := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		rst = append(rst, n.Val)
		getAllElementsGoLeft(n.Right, s)
	}

	return rst
}

func getAllElementsGoLeft(node *TreeNode, stack *[]*TreeNode) {
	for node != nil {
		*stack = append(*stack, node)
		node = node.Left
	}
}
