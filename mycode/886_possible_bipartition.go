package mycode

func possibleBipartitionDfs(node int, graph [][]int, curFlag bool, flagMap map[int]bool) bool {
	if f, ok := flagMap[node]; ok {
		return f == curFlag
	} else {
		flagMap[node] = curFlag
	}

	for _, neighbor := range graph[node] {
		if !possibleBipartitionDfs(neighbor, graph, !curFlag, flagMap) {
			return false
		}
	}
	return true
}

func possibleBipartition(N int, dislikes [][]int) bool {
	// store the graph in an edge-list style
	graph := make([][]int, N+1)
	for i, _ := range graph {
		graph[i] = make([]int, 0)
	}

	// build graph
	for _, v := range dislikes {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	flagMap := make(map[int]bool)
	for i := 1; i <= N; i++ {
		if _, ok := flagMap[i]; !ok && !possibleBipartitionDfs(i, graph, false, flagMap) {
			return false
		}
	}
	return true
}
