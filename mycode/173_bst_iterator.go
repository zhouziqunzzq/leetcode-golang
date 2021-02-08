package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/binary-search-tree-iterator/solution/
// Solution 2: iterative approach using stack
type BSTIterator struct {
	s []*TreeNode
}

func Constructor173(root *TreeNode) BSTIterator {
	it := BSTIterator{
		s: make([]*TreeNode, 0),
	}
	it.goLeft(root)
	return it
}

func (this *BSTIterator) goLeft(n *TreeNode) {
	for n != nil {
		this.s = append(this.s, n)
		n = n.Left
	}
}

func (this *BSTIterator) Next() int {
	n := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]

	if n.Right != nil {
		this.goLeft(n.Right)
	}

	return n.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.s) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
