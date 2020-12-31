package mycode

// Solution 1: monotonically increasing stack
// https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/28917/AC-Python-clean-solution-using-stack-76ms
// Solution 2: dp-style with two auxiliary arrays
// https://leetcode.com/problems/largest-rectangle-in-histogram/discus/28902/5ms-O(n)-Java-solution-explained-(beats-96)
// We use solution 1 here.
func largestRectangleArea(heights []int) int {
	// append sentinel building
	heights = append(heights, 0)
	n := len(heights)

	s := make([]int, 1)
	s[0] = 0
	ans := heights[0]

	for i := 1; i < n; i++ {
		// pop all building idx higher than current heights[i]
		for len(s) > 0 && heights[s[len(s)-1]] > heights[i] {
			curIdx := s[len(s)-1]
			s = s[:len(s)-1]
			h := heights[curIdx]
			// the width calculation is the tricky part
			// 1. s is empty: all previous heights > heights[curIdx]
			//    so we can use i as the width
			// 2. s isn't empty: the top is the first one to the left less than heights[curIdx]
			//    then we use i - (s.top() + 1) as the width
			w := i
			if len(s) > 0 {
				w = i - s[len(s)-1] - 1
			}
			ans = maxInt(ans, h*w)
		}
		s = append(s, i)
	}

	// restore original inputs
	heights = heights[:n-1]

	return ans
}
