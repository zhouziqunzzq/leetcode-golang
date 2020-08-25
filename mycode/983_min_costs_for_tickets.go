package mycode

import "math"

// opt[i] = min cost considering days 0...i-1
// opt[i] = min(
//     opt[i-1] + costs[0]
//     min(opt[j]) + costs[1] where days[i] - days[j] < 7
//     min(opt[j]) + costs[2] where days[i] - days[j] < 30
// )
func mincostTickets(days []int, costs []int) int {
	opt := make([]int, len(days)+1)

	// base case
	opt[0] = 0
	opt[1] = minInt3(costs[0], costs[1], costs[2])

	// fill the table
	for i := 2; i <= len(days); i++ {
		opt[i] = math.MaxInt32
		// selection 1
		opt[i] = minInt(opt[i], opt[i-1]+costs[0])
		// selection 2
		var j int
		for j = i - 1; j > 0 && days[i-1]-days[j-1] < 7; j-- {
			opt[i] = minInt(opt[i], opt[j]+costs[1])
		}
		if j >= 0 {
			opt[i] = minInt(opt[i], opt[j]+costs[1])
		}
		// selection 3
		for j = i - 1; j > 0 && days[i-1]-days[j-1] < 30; j-- {
			opt[i] = minInt(opt[i], opt[j]+costs[2])
		}
		if j >= 0 {
			opt[i] = minInt(opt[i], opt[j]+costs[2])
		}
	}

	return opt[len(days)]
}
