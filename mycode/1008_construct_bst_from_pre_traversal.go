package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	lastNode := root
	// s act as a monotonously decreasing stack
	s := make([]*TreeNode, 0)
	s = append(s, root)

	for _, nowVal := range preorder[1:] {
		if nowVal < s[len(s)-1].Val {
			// going left as deep as possible
			nowNode := &TreeNode{nowVal, nil, nil}
			lastNode.Left = nowNode
			s = append(s, nowNode)
			lastNode = nowNode
		} else {
			// can't go left, need to build tree and try to go right
			for len(s) > 0 && s[len(s)-1].Val < nowVal {
				nowNode := s[len(s)-1]
				s = s[:len(s)-1]
				lastNode = nowNode
			}
			lastNode.Right = &TreeNode{nowVal, nil, nil}
			s = append(s, lastNode.Right)
			lastNode = s[len(s)-1]
		}
	}

	return root
}
