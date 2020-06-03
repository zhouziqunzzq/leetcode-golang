package mycode

import "sort"

// Solution 1: Greedy. Sort by diff.
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	ans := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			ans += costs[i][0]
		} else {
			ans += costs[i][1]
		}
	}

	return ans
}

// Solution 2: Dp
func twoCitySchedCostDp(costs [][]int) int {
	// opt[i][j]: min cost considering costs[0...i+j]
	// where i people go to A and j people go to B
	N := len(costs) / 2
	var opt [51][51]int

	// base cases
	for i := 1; i <= N; i++ {
		opt[i][0] = opt[i-1][0] + costs[i-1][0]
	}
	for j := 1; j <= N; j++ {
		opt[0][j] = opt[0][j-1] + costs[j-1][1]
	}

	// fill the table
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			opt[i][j] = minInt(opt[i-1][j]+costs[i+j-1][0],
				opt[i][j-1]+costs[i+j-1][1])
		}
	}

	return opt[N][N]
}
