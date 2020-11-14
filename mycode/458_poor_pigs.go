package mycode

import "math"

// https://leetcode.com/problems/poor-pigs/discuss/94273/Solution-with-detailed-explanation
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	rounds := minutesToTest / minutesToDie
	base := rounds + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(base))))
}
