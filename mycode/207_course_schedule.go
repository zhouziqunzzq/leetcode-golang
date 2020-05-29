package mycode

// flag: 0 - not visited, 1 - visiting, 2 - visited
func canFinishDfs(node int, graph [][]int, flags []uint8) bool {
	if flags[node] == 1 {
		// circle!
		return true
	}
	if flags[node] == 2 {
		return false
	}

	// start visiting this node
	flags[node] = 1
	for _, neighbor := range graph[node] {
		if canFinishDfs(neighbor, graph, flags) {
			return true
		}
	}

	// finish visiting this node
	flags[node] = 2
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i, _ := range graph {
		graph[i] = make([]int, 0)
	}
	for _, e := range prerequisites {
		graph[e[0]] = append(graph[e[0]], e[1])
	}
	flags := make([]uint8, numCourses)

	for i := 0; i < numCourses; i++ {
		if canFinishDfs(i, graph, flags) {
			// has circle!
			return false
		}
	}
	return true
}
