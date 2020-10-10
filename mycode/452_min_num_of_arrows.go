package mycode

import "sort"

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/discuss/93703/Share-my-explained-Greedy-solution
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	// sort w.r.t right end in ascending order
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	ans := 0
	i := 0
	for ; i < len(points); {
		// greedy: shoot at right end
		r := points[i][1]
		ans++
		i++
		for i < len(points) && points[i][0] <= r {
			i++
		}
	}

	return ans
}
