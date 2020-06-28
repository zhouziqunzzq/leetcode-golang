package mycode

import "container/heap"

func findItineraryHelper(now string, g map[string]*StringPriorityQueue, rst *[]string) {
	if pq, ok := g[now]; ok {
		for pq.Len() > 0 {
			next := heap.Pop(pq).(string)
			// DFS
			findItineraryHelper(next, g, rst)
		}
	}

	*rst = append(*rst, now)
}

func findItinerary(tickets [][]string) []string {
	g := make(map[string]*StringPriorityQueue)
	rst := make([]string, 0)

	for _, t := range tickets {
		if pq, ok := g[t[0]]; ok {
			heap.Push(pq, t[1])
		} else {
			newPq := make(StringPriorityQueue, 0)
			g[t[0]] = &newPq
			heap.Push(&newPq, t[1])
		}
	}

	findItineraryHelper("JFK", g, &rst)

	reverseSlice(rst)

	return rst
}
