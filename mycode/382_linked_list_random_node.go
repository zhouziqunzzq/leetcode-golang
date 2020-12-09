package mycode

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/linked-list-random-node/solution/
// Solution 2: Reservoir algorithm
// https://en.wikipedia.org/wiki/Reservoir_sampling
type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor382(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	i := 1
	rst := 0
	cur := this.head

	for cur != nil {
		if rand.Float64() < 1.0/float64(i) {
			rst = cur.Val
		}
		i++
		cur = cur.Next
	}

	return rst
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
