package mycode

// https://leetcode.com/problems/critical-connections-in-a-network/discuss/382638/DFS-detailed-explanation-O(orEor)-solution
// DFS
// Alternative: tarjan's algorithm (to find strongly connected components)
func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	connMap := make(map[int]map[int]bool)
	for _, c := range connections {
		a, b := c[0], c[1]

		if g[a] == nil {
			g[a] = make([]int, 0)
		}
		g[a] = append(g[a], b)
		if g[b] == nil {
			g[b] = make([]int, 0)
		}
		g[b] = append(g[b], a)

		if _, ok := connMap[a]; !ok {
			connMap[a] = make(map[int]bool)
		}
		connMap[a][b] = true
	}

	rank := make([]int, n)
	for i := range rank {
		rank[i] = -2
	}

	criticalConnectionsDFS(0, 0, g, rank, connMap)

	rst := make([][]int, 0)
	for a, m := range connMap {
		for b := range m {
			t := []int{a, b}
			rst = append(rst, t)
		}
	}

	return rst
}

// return the minimum rank encountered during visiting neighbors
func criticalConnectionsDFS(
	node int, depth int, g [][]int, rank []int, connMap map[int]map[int]bool) int {
	if rank[node] >= 0 {
		// visiting or visited
		return rank[node]
	}

	n := len(rank)
	rank[node] = depth
	minRank := n

	for _, neighbor := range g[node] {
		if rank[neighbor] == depth-1 {
			continue // prevent from going back to immediate parent
		}
		nextRank := criticalConnectionsDFS(neighbor, depth+1, g, rank, connMap)
		if nextRank <= depth {
			// (node, neighbor) is on a cycle
			criticalConnectionsDeleteConnection(node, neighbor, connMap)
			criticalConnectionsDeleteConnection(neighbor, node, connMap)
		}
		minRank = minInt(minRank, nextRank)
	}

	rank[node] = n
	return minRank
}

func criticalConnectionsDeleteConnection(a, b int, connMap map[int]map[int]bool) {
	if m, ok := connMap[a]; ok {
		if _, ok1 := m[b]; ok1 {
			delete(m, b)
		}
		if len(m) == 0 {
			delete(connMap, a)
		}
	}
}
