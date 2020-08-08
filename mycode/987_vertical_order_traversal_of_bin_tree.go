package mycode

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node987 struct {
	n *TreeNode
	x int
	y int
}

func verticalTraversal(root *TreeNode) [][]int {
	rst := make([][]int, 0)
	m := make(map[int]map[int][]int)
	// BFS queue
	q := make([]*Node987, 1)
	q[0] = &Node987{root, 0, 0}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		// visit current node
		if _, ok := m[n.x]; !ok {
			m[n.x] = make(map[int][]int)
		}
		if _, ok := m[n.x][n.y]; !ok {
			m[n.x][n.y] = make([]int, 0)
		}
		m[n.x][n.y] = append(m[n.x][n.y], n.n.Val)
		// visit left child
		if n.n.Left != nil {
			q = append(q, &Node987{n.n.Left, n.x - 1, n.y + 1})
		}
		// visit right child
		if n.n.Right != nil {
			q = append(q, &Node987{n.n.Right, n.x + 1, n.y + 1})
		}
	}

	xs := make([]int, len(m))
	i := 0
	for x := range m {
		xs[i] = x
		i++
	}
	sort.Ints(xs)

	for _, x := range xs {
		rst = append(rst, make([]int, 0))
		ys := make([]int, len(m[x]))
		i := 0
		for y := range m[x] {
			ys[i] = y
			i++
		}
		sort.Ints(ys)
		for _, y := range ys {
			ts := m[x][y]
			sort.Ints(ts)
			for _, t := range ts {
				rst[len(rst)-1] = append(rst[len(rst)-1], t)
			}
		}
	}

	return rst
}
