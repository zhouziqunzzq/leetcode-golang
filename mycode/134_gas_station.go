package mycode

// https://leetcode.com/problems/gas-station/discuss/42568/Share-some-of-my-ideas.
func canCompleteCircuit(gas []int, cost []int) int {
	tank := 0
	start := 0
	sumGas, sumCost := 0, 0

	for i := 0; i < len(gas); i++ {
		sumGas += gas[i]
		sumCost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if sumGas >= sumCost {
		return start
	} else {
		return -1
	}
}
