package mycode

import "strings"
import "math"

// https://leetcode.com/problems/find-the-shortest-superstring/discuss/194932/Travelling-Salesman-Problem
// OH MY GOD this is an NP-hard TSP problem underneath...
// Here we offer a DP solution to TSP...
func shortestSuperstring(words []string) string {
	n := len(words)
	// graph[i][j] means the length of string to append when A[i] followed by A[j].
	// eg. A[i] = abcd, A[j] = bcde, then graph[i][j] = 1
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	// build the graph
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = shortestSuperstringHelper(words[i], words[j])
			graph[j][i] = shortestSuperstringHelper(words[j], words[i])
		}
	}
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	path := make([][]int, 1<<n)
	for i := range path {
		path[i] = make([]int, n)
	}
	last := -1
	minPathLen := math.MaxInt32

	// start TSP DP
	for i := 1; i < (1 << n); i++ { // i - bitset(mask) of nodes
		fillIntSlice(dp[i], math.MaxInt32)
		for j := 0; j < n; j++ { // j - TSP path ending at node j
			if i&(1<<j) > 0 {
				prev := i ^ (1 << j)
				if prev == 0 { // base case
					dp[i][j] = len(words[j])
				} else {
					// dp[i][j] = max(dp[prev][k] + graph[k][j]) for k = 1...n
					for k := 0; k < n; k++ {
						if dp[prev][k] < math.MaxInt32 && dp[prev][k]+graph[k][j] < dp[i][j] {
							dp[i][j] = dp[prev][k] + graph[k][j]
							path[i][j] = k
						}
					}
				}
			}
			if i == (1<<n)-1 && dp[i][j] < minPathLen { // update global minimum
				minPathLen = dp[i][j]
				last = j
			}
		}
	}

	// backtrace using path
	sb := strings.Builder{}
	cur := 1<<n - 1
	stk := make([]int, 0)
	for cur > 0 {
		stk = append(stk, last)
		tCur := cur
		cur ^= 1 << last
		last = path[tCur][last]
	}

	// build the result
	i := stk[len(stk)-1]
	sb.WriteString(words[i])
	stk = stk[:len(stk)-1]
	for len(stk) > 0 {
		j := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		sb.WriteString(words[j][len(words[j])-graph[i][j]:])
		i = j
	}

	return sb.String()
}

func shortestSuperstringHelper(a, b string) int {
	for i := 1; i < len(a); i++ {
		if strings.HasPrefix(b, a[i:]) {
			return len(b) - (len(a) - i)
		}
	}

	return len(b)
}
