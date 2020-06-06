package mycode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// sort by height([0]) and k([1])
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] ||
			(people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	for i := 1; i < len(people); i++ {
		if people[i][1] != i {
			targetPos := people[i][1]
			for j := i; j > targetPos; j-- {
				people[j-1], people[j] = people[j], people[j-1]
			}
		}
	}

	return people
}
