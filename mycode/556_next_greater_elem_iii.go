package mycode

import "strconv"
import "sort"
import "math"

// https://leetcode.com/problems/next-greater-element-iii/discuss/101824/Simple-Java-solution-(4ms)-with-explanation.
// same as next permutation
func nextGreaterElement(n int) int {
	ss := []rune(strconv.Itoa(n))
	sn := len(ss)
	i, j := 0, 0

	for i = sn - 1; i > 0; i-- {
		if ss[i-1] < ss[i] {
			break
		}
	}
	if i == 0 {
		return -1
	}

	x := ss[i-1]
	minPtr := i
	for j = i + 1; j < sn; j++ {
		if ss[j] > x && ss[j] <= ss[minPtr] {
			minPtr = j
		}
	}
	ss[i-1], ss[minPtr] = ss[minPtr], ss[i-1]

	trailing := ss[i:]
	sort.Slice(trailing, func(i, j int) bool {
		return trailing[i] < trailing[j]
	})

	ans, _ := strconv.Atoi(string(ss))
	if ans <= math.MaxInt32 {
		return ans
	} else {
		return -1
	}
}
