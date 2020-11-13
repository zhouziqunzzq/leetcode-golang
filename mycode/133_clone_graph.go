package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return nil
	}

	nodes := make([]*Node133, 110)
	return cloneGraphHelper(node, nodes)
}

func cloneGraphHelper(node *Node133, nodes []*Node133) *Node133 {
	id := node.Val
	if nodes[id] == nil {
		nodes[id] = &Node133{
			Val:       id,
			Neighbors: make([]*Node133, len(node.Neighbors)),
		}
		for i, n := range node.Neighbors {
			nodes[id].Neighbors[i] = cloneGraphHelper(n, nodes)
		}
	}
	return nodes[id]
}
