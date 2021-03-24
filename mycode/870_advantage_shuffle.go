package mycode

import "sort"

// https://leetcode.com/problems/advantage-shuffle/solution/
// Greedy
func advantageCount(A []int, B []int) []int {
	sortedA := A
	sort.Ints(sortedA)
	sortedB := make([]int, len(B))
	for j, b := range B {
		sortedB[j] = b
	}
	sort.Ints(sortedB)

	arranged := make(map[int][]int)
	remained := make([]int, 0, len(A))

	// greedily select cards from A that can beat cards in B
	j := 0
	for _, a := range sortedA {
		b := sortedB[j]
		if a > b {
			if _, ok := arranged[b]; !ok {
				arranged[b] = make([]int, 0)
			}
			arranged[b] = append(arranged[b], a) // we can use this a to beat b...
			j++                                  // and proceed to next b in B
		} else {
			remained = append(remained, a)
		}
	}
	// at the end of previous loop, every a from A is either
	// in arranged, where it's used to beat b, or it's in remained

	rst := make([]int, len(A))
	for i := range rst {
		b := B[i]
		if al, ok := arranged[b]; ok && len(al) > 0 {
			rst[i] = al[0]
			arranged[b] = arranged[b][1:]
		} else {
			rst[i] = remained[0]
			remained = remained[1:]
		}
	}

	return rst
}
