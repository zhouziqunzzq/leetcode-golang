package mycode

import "sort"

// https://leetcode.com/problems/boats-to-save-people/solution/
// Greedy
func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] < people[j]
	})

	i, j := 0, len(people)-1
	ans := 0

	for i <= j {
		// put the heaviest person people[j] into a new boat...
		ans++
		// ...and try to put the lightest person people[i] into the same boat
		if people[i]+people[j] <= limit {
			i++
		}
		j--
	}

	return ans
}
