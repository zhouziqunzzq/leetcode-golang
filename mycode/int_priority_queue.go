package mycode

import "container/heap"

// A PriorityQueue implements heap.Interface and holds integers.

type IntItem struct {
	Val int
	Idx int
}

type IntPriorityQueue struct {
	items     []*IntItem
	itemMap   map[int][]*IntItem
	IsMaxHeap bool
}

func CreateIntPriorityQueue(isMaxHeap bool) *IntPriorityQueue {
	pq := &IntPriorityQueue{
		items:     make([]*IntItem, 0),
		itemMap:   make(map[int][]*IntItem),
		IsMaxHeap: isMaxHeap,
	}
	heap.Init(pq)
	return pq
}

func (pq IntPriorityQueue) Len() int { return len(pq.items) }

func (pq IntPriorityQueue) Less(i, j int) bool {
	if pq.IsMaxHeap {
		return pq.items[i].Val > pq.items[j].Val
	} else {
		return pq.items[i].Val < pq.items[j].Val
	}
}

func (pq IntPriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Idx = i
	pq.items[j].Idx = j
}

func (pq IntPriorityQueue) addItemEntry(item *IntItem) {
	val := item.Val
	if _, ok := pq.itemMap[val]; !ok {
		// init entry
		pq.itemMap[val] = make([]*IntItem, 1)
		pq.itemMap[val][0] = item
	} else {
		// append entry
		pq.itemMap[val] = append(pq.itemMap[val], item)
	}
}

func (pq IntPriorityQueue) getItemEntry(val int) *IntItem {
	if items, ok := pq.itemMap[val]; ok && len(items) > 0 {
		return pq.itemMap[val][0]
	} else {
		return nil
	}
}

func (pq IntPriorityQueue) removeItemEntry(val int) *IntItem {
	if items, ok := pq.itemMap[val]; ok && len(items) > 0 {
		item := pq.itemMap[val][0]
		pq.itemMap[val] = pq.itemMap[val][1:]
		return item
	} else {
		return nil
	}
}

func (pq *IntPriorityQueue) Push(x interface{}) {
	n := pq.Len()
	item := x.(*IntItem)
	item.Idx = n
	pq.items = append(pq.items, item)
	pq.addItemEntry(item)
}

func (pq *IntPriorityQueue) PushInt(val int) {
	heap.Push(pq, &IntItem{
		Val: val,
		Idx: -1,
	})
}

func (pq *IntPriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	lastItem := old[n-1]
	selectedItem := pq.removeItemEntry(lastItem.Val)
	pq.Swap(lastItem.Idx, selectedItem.Idx)
	old[n-1] = nil        // avoid memory leak
	selectedItem.Idx = -1 // for safety
	pq.items = old[0 : n-1]
	return selectedItem
}

func (pq *IntPriorityQueue) PopInt() int {
	return heap.Pop(pq).(*IntItem).Val
}

func (pq *IntPriorityQueue) RemoveInt(val int) {
	item := pq.getItemEntry(val)
	if item != nil {
		heap.Remove(pq, item.Idx)
	}
}

func (pq *IntPriorityQueue) PeekInt() int {
	return pq.items[0].Val
}
