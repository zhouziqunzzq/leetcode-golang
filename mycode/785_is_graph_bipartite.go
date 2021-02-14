package mycode

// https://leetcode.com/problems/is-graph-bipartite/discuss/115487/Java-Clean-DFS-solution-with-Explanation
// The gist is trying to color the graph using two colors (DFS/BFS)
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))

	for i := range graph {
		if colors[i] == 0 && !isBipartiteDFS(graph, colors, 1, i) {
			return false
		}
	}

	return true
}

func isBipartiteDFS(graph [][]int, colors []int, color, node int) bool {
	if colors[node] != 0 {
		return colors[node] == color
	}

	colors[node] = color
	for _, next := range graph[node] {
		if !isBipartiteDFS(graph, colors, -color, next) {
			return false
		}
	}

	return true
}
