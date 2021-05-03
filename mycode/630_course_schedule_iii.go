package mycode

import "sort"

// https://leetcode.com/problems/course-schedule-iii/solution/
// Solution 4: optimized iterative greedy approach
// Greedy: 1. Sort by end date.
// 		   2. Swap with the previous selected course that has the largest duration.
//		      (Pick the course with smaller duration).
func scheduleCourse(courses [][]int) int {
	// Greedy 1
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	time := 0
	cnt := 0
	for i, c := range courses {
		if time+c[0] <= c[1] {
			// Greedy 1: take this course
			time += c[0]
			courses[cnt] = c // maintain that courses[0...cnt] holds chosen courses so far
			cnt++
		} else {
			// Greedy 2: try to swap with a previously taken course
			maxIdx := i
			for j := 0; j < cnt; j++ {
				if courses[j][0] > courses[maxIdx][0] {
					maxIdx = j
				}
			}
			if courses[maxIdx][0] > c[0] {
				time += c[0] - courses[maxIdx][0]
				courses[maxIdx] = c
			}
		}
	}

	return cnt
}
