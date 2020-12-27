package mycode

// https://leetcode.com/problems/jump-game-iv/solution/
// BFS
func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	// elem val -> list of node ids (elem idx)
	graph := make(map[int][]int)
	for i, v := range arr {
		if _, ok := graph[v]; !ok {
			graph[v] = make([]int, 0)
		}
		graph[v] = append(graph[v], i)
	}

	visited := make(map[int]bool)
	cur := make([]int, 1)
	cur[0] = 0
	step := 0

	for len(cur) > 0 {
		next := make([]int, 0)

		for _, node := range cur {
			// check target
			if node == n-1 {
				return step
			}

			// expand children with same val
			for _, child := range graph[arr[node]] {
				if ok := visited[child]; !ok && child != node {
					visited[child] = true
					next = append(next, child)
				}
			}

			// remove entry of arr[node] from graph to avoid stepping back
			graph[arr[node]] = nil

			// expand neighbors
			if _, ok := visited[node-1]; !ok && node-1 >= 0 {
				visited[node-1] = true
				next = append(next, node-1)
			}
			if _, ok := visited[node+1]; !ok && node+1 < n {
				visited[node+1] = true
				next = append(next, node+1)
			}
		}

		cur = next
		step++
	}

	return -1
}
