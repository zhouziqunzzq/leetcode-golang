package mycode

// https://leetcode.com/problems/sliding-window-maximum/discuss/65884/Java-O(n)-solution-using-deque-with-explanation
// Solution using a deque (double ends queue)
// Actually it's a monotonically decreasing deque
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	q := make([]int, 0, k)
	ans := make([]int, n-k+1)
	ri := 0

	for i, cur := range nums {
		// remove indexes out of range of [i-k+1, i] from front
		if len(q) > 0 && q[0] < i-k+1 {
			q = q[1:]
		}
		// remove indexes pointing to elem less than current number from end
		for len(q) > 0 && nums[q[len(q)-1]] < cur {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			ans[ri] = nums[q[0]]
			ri++
		}
	}

	return ans
}
