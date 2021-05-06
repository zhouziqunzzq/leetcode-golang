package mycode

// https://leetcode.com/problems/jump-game-ii/discuss/18028/O(n)-BFS-solution
// BFS inspired Greedy algo
// Greedy: jump as far as one can at every level
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	level, curMax, nextMax := 0, 0, 0
	i := 0

	for curMax-i+1 > 0 { // while we haven't run out of nodes to visit
		level++

		for ; i <= curMax; i++ {
			nextMax = maxInt(nextMax, i+nums[i])
			if nextMax >= len(nums)-1 { // end-of-array reachable
				return level
			}
		}

		curMax = nextMax
	}

	return 0
}
