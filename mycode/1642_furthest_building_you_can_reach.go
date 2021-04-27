package mycode

// https://leetcode.com/problems/furthest-building-you-can-reach/discuss/918515/JavaC%2B%2BPython-Priority-Queue
// Greedy. Use a priority queue to store the height diffs, and use ladders for larger diffs whenever possible.
func furthestBuilding(heights []int, bricks int, ladders int) int {
	pq := CreateIntPriorityQueue(false)

	for i := 0; i < len(heights)-1; i++ {
		d := heights[i+1] - heights[i]
		if d > 0 {
			pq.PushInt(d)
		}
		if pq.Len() > ladders {
			bricks -= pq.PopInt()
		}
		if bricks < 0 {
			return i
		}
	}

	return len(heights) - 1
}
