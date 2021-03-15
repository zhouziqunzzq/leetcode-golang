package mycode

import "strconv"
import "strings"
import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/serialize-and-deserialize-bst/discuss/177617/the-General-Solution-for-Serialize-and-Deserialize-BST-and-Serialize-and-Deserialize-BT
type Codec449 struct {
}

func Constructor449() Codec449 {
	return Codec449{}
}

// Serializes a tree to a single string.
func (this *Codec449) serialize(root *TreeNode) string {
	// serialize using pre-order traversal
	sb := strings.Builder{}
	this.serializeHelper(root, &sb)
	n := sb.Len()
	if n == 0 {
		return sb.String()
	} else {
		return sb.String()[:n-1]
	}
}

func (this *Codec449) serializeHelper(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		return
	}
	sb.WriteString(strconv.Itoa(node.Val))
	sb.WriteRune(',')
	this.serializeHelper(node.Left, sb)
	this.serializeHelper(node.Right, sb)
}

// Deserializes your encoded data to tree.
func (this *Codec449) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	s := strings.Split(data, ",")
	p := 0
	return this.deserializeHelper(s, &p, math.MinInt32, math.MaxInt32)
}

func (this *Codec449) deserializeHelper(s []string, p *int, lowerBound, upperBound int) *TreeNode {
	if *p >= len(s) {
		return nil
	}

	curVal, _ := strconv.Atoi(s[*p])
	if curVal < lowerBound || curVal > upperBound {
		return nil
	}

	*p++
	return &TreeNode{
		Val:   curVal,
		Left:  this.deserializeHelper(s, p, lowerBound, curVal),
		Right: this.deserializeHelper(s, p, curVal, upperBound),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
