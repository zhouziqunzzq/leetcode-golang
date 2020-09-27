package mycode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// graph["a"]["b"] = value of a/b
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(values); i++ {
		a, b, val := equations[i][0], equations[i][1], values[i]
		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = val
		graph[b][a] = 1.0 / val
	}

	visited := make(map[string]bool)
	ans := make([]float64, len(queries))
	for i, q := range queries {
		// sanity checks
		if _, ok := graph[q[0]]; !ok {
			ans[i] = -1.0
			continue
		}
		if _, ok := graph[q[1]]; !ok {
			ans[i] = -1.0
			continue
		}

		// evaluate
		rst, isSuccess := calcEquationDFS(
			graph, visited,
			q[0], q[1],
			1.0,
		)
		if isSuccess {
			ans[i] = rst
		} else {
			ans[i] = -1.0
		}
	}

	return ans
}

func calcEquationDFS(
	graph map[string]map[string]float64,
	visited map[string]bool,
	cur, target string,
	curVal float64,
) (float64, bool) {
	if cur == target {
		return curVal, true
	}

	visited[cur] = true
	rst := -1.0
	isSuccess := false
	for next, nextVal := range graph[cur] {
		if nextVisited, ok := visited[next]; !ok || !nextVisited {
			rst, isSuccess = calcEquationDFS(graph, visited, next, target, curVal*nextVal)
			if isSuccess {
				break
			}
		}
	}
	visited[cur] = false
	return rst, isSuccess
}
