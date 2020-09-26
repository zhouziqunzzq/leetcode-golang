package mycode

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	ans := 0
	due := timeSeries[0] + duration - 1
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] > due {
			ans += duration
		} else {
			ans += timeSeries[i] - timeSeries[i-1]
		}
		due = timeSeries[i] + duration - 1
	}
	ans += duration
	return ans
}
