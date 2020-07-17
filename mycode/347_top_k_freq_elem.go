package mycode

import "math/rand"

// Solution: quick select https://leetcode.com/articles/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		if _, ok := count[n]; ok {
			count[n]++
		} else {
			count[n] = 1
		}
	}

	unique := make([]int, len(count))
	i := 0
	for k, _ := range count {
		unique[i] = k
		i++
	}

	topKFrequentQuickSelect(unique, count, 0, len(unique)-1, len(unique)-k)

	return unique[len(unique)-k:]
}

// l: left, r: right, p: pivot
func topKFrequentPartition(target []int, count map[int]int, l, r, p int) int {
	pivotFreq := count[target[p]]

	target[r], target[p] = target[p], target[r]
	storeIdx := l

	for i := l; i <= r; i++ {
		if count[target[i]] < pivotFreq {
			target[i], target[storeIdx] = target[storeIdx], target[i]
			storeIdx++
		}
	}

	target[storeIdx], target[r] = target[r], target[storeIdx]

	return storeIdx
}

func topKFrequentQuickSelect(target []int, count map[int]int, l, r, kSmallest int) {
	if l == r {
		return
	}

	p := l + rand.Intn(r-l+1)

	pivotIdx := topKFrequentPartition(target, count, l, r, p)

	if pivotIdx == kSmallest {
		return
	} else if pivotIdx < kSmallest {
		topKFrequentQuickSelect(target, count, pivotIdx+1, r, kSmallest)
	} else {
		topKFrequentQuickSelect(target, count, l, pivotIdx-1, kSmallest)
	}
}
