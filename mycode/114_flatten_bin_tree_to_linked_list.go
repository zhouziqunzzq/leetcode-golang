package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/37010/Share-my-simple-NON-recursive-solution-O(1)-space-complexity!
// Non-recursive solution
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// pre points to the pre-node(in terms of pre-order traversal) of the cur node
			// which should link to the right subtree of cur node.
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right

			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
