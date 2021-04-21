package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type NaryNode struct {
	Val      int
	Children []*NaryNode
}

func preorder(root *NaryNode) []int {
	rst := make([]int, 0)
	preorderHelper(root, &rst)
	return rst
}

func preorderHelper(n *NaryNode, rst *[]int) {
	if n == nil {
		return
	}

	*rst = append(*rst, n.Val)
	for _, c := range n.Children {
		preorderHelper(c, rst)
	}
}
