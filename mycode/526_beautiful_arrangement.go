package mycode

// https://leetcode.com/problems/beautiful-arrangement/solution/
// Solution 3: DFS + backtracking
func countArrangement(n int) int {
	visited := make([]bool, n+1)
	rst := 0
	countArrangementHelper(n, 1, visited, &rst)
	return rst
}

func countArrangementHelper(n, pos int, visited []bool, rst *int) {
	if pos > n {
		*rst++
		return
	}

	for i := 1; i <= n; i++ {
		if !visited[i] && (i%pos == 0 || pos%i == 0) {
			visited[i] = true
			countArrangementHelper(n, pos+1, visited, rst)
			visited[i] = false
		}
	}
}
