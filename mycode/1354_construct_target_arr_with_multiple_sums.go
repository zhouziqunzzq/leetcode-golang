package mycode

// https://leetcode.com/problems/construct-target-array-with-multiple-sums/discuss/510256/JavaC%2B%2BPython-Backtrack-OJ-is-wrong
// Greedy + reverse the problem
func isPossible(target []int) bool {
	total := 0
	for _, t := range target {
		total += t
	}

	pq := FasterPriorityQueue{}
	pq.Init(target, func(a pqElemType, b pqElemType) bool {
		return a > b // max heap
	})

	for {
		a := *(pq.Pop()) // a is the current max elem
		total -= a       // [a, ...total-a...]

		if a == 1 || total == 1 {
			// check if total==1 to handle a corner case "[10000, 1]" later when we check if a%total==0
			return true
		}
		if a < total || total == 0 || a%total == 0 {
			return false
		}

		a %= total // use % here to accelerate the backtracking process
		total += a
		pq.Push(a)
	}
}
