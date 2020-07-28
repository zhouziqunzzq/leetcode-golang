package mycode

// https://leetcode.com/problems/task-scheduler/discuss/104500/Java-O(n)-time-O(1)-space-1-pass-no-sorting-solution-with-detailed-explanation
func leastInterval(tasks []byte, n int) int {
	// greedy: arrange the most frequent jobs first
	var cnt [26]int
	max := 0
	maxCnt := 0

	for _, t := range tasks {
		idx := t - 'A'
		cnt[idx]++
		if cnt[idx] > max {
			max = cnt[idx]
			maxCnt = 1
		} else if cnt[idx] == max {
			maxCnt++
		}
	}

	idlePartCnt := max - 1
	idlePartLen := n - (maxCnt - 1)
	emptySlots := idlePartCnt * idlePartLen
	trivialTasks := len(tasks) - maxCnt*max
	idles := maxInt(0, emptySlots-trivialTasks)

	return len(tasks) + idles
}
