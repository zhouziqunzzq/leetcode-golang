package mycode

// A StringPriorityQueue implements heap.Interface and holds strings.
type StringPriorityQueue []string

func (pq StringPriorityQueue) Len() int { return len(pq) }

func (pq StringPriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq StringPriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *StringPriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(string)) }

func (pq *StringPriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	*pq = old[0 : len(old)-1]
	return item
}
