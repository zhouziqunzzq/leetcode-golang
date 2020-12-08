package mycode

// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/solution/
func numPairsDivisibleBy60(time []int) int {
	remainders := make([]int, 60)
	ans := 0

	for _, t := range time {
		if t%60 == 0 {
			// a % 60 == 0 && b % 60 == 0
			ans += remainders[0]
		} else {
			// a % 60 + b % 60 == 60
			ans += remainders[60-t%60]
		}
		remainders[t%60]++
	}

	return ans
}
