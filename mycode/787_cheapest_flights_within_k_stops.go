package mycode

import "math"

// Run Bellman-Ford K+1 times
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	opt := make([]int, n)
	for i, _ := range opt {
		opt[i] = math.MaxInt32
	}
	opt[src] = 0

	for k := 0; k < K + 1; k++ {
		optNew := make([]int, n)
		for i, v := range opt {
			optNew[i] = v
		}
		for _, e := range flights {
			optNew[e[1]] = minInt(optNew[e[1]], opt[e[0]] + e[2])
		}
		opt = optNew
	}

	if opt[dst] < math.MaxInt32 {
		return opt[dst]
	} else {
		return -1
	}
}
