package mycode

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i, _ := range graph {
		graph[i] = make([]int, 0)
	}
	for _, e := range prerequisites {
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	flags := make([]uint8, numCourses)

	rst := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if flags[i] == 0 {
			if !findOrderDFS(i, graph, flags, &rst) {
				return make([]int, 0)
			}
		}
	}

	reverseSlice(rst)
	return rst
}

// flag: 0 - not visited, 1 - visiting, 2 - visited
func findOrderDFS(node int, graph [][]int, flags []uint8, rst *[]int) bool {
	// start visiting this node
	flags[node] = 1
	canSort := true
	for _, neighbor := range graph[node] {
		if flags[neighbor] == 1 {
			// cycle detected
			canSort = false
			break
		} else if flags[neighbor] == 0 {
			canSort = canSort && findOrderDFS(neighbor, graph, flags, rst)
			if !canSort {
				// early stopping
				break
			}
		}
	}

	// finish visiting this node
	flags[node] = 2
	*rst = append(*rst, node)
	return canSort
}
