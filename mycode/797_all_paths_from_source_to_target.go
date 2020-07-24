package mycode

func allPathsSourceTarget(graph [][]int) [][]int {
	rst := make([][]int, 0)
	path := make([]int, 1)

	path[0] = 0
	allPathsSourceTargetDFS(graph, 0, &rst, &path)

	return rst
}

func allPathsSourceTargetDFS(graph [][]int, node int, rst *[][]int, path *[]int) {
	if node == len(graph)-1 {
		*rst = append(*rst, make([]int, len(*path)))
		copy((*rst)[len(*rst)-1], *path)
		return
	}

	for _, next := range graph[node] {
		*path = append(*path, next)
		allPathsSourceTargetDFS(graph, next, rst, path)
		*path = (*path)[:len(*path)-1]
	}
}
