package mycode

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])
	if n == 0 {
		return n
	}

	// 2-d dp array -> 1-d dp array
	opt := make([]int, n+1)
	for i, _ := range opt {
		opt[i] = math.MaxInt32
	}
	opt[n-1] = 1

	// fill from bottom-right to top-left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			opt[j] = minInt(opt[j], opt[j+1]) - dungeon[i][j]
			// don't waste hp :p so use max to cap it
			opt[j] = maxInt(1, opt[j])
		}
	}

	return opt[0]
}
