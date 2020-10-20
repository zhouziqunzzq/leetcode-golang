package mycode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodes := make([]*Node, 110)
	return cloneGraphHelper(node, nodes)
}

func cloneGraphHelper(node *Node, nodes []*Node) *Node {
	id := node.Val
	if nodes[id] == nil {
		nodes[id] = &Node{
			Val:       id,
			Neighbors: make([]*Node, len(node.Neighbors)),
		}
		for i, n := range node.Neighbors {
			nodes[id].Neighbors[i] = cloneGraphHelper(n, nodes)
		}
	}
	return nodes[id]
}
