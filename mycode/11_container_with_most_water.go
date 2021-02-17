package mycode

// https://leetcode.com/problems/container-with-most-water/discuss/6090/Simple-and-fast-C%2B%2BC-with-explanation
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0

	for l < r {
		h := minInt(height[l], height[r])
		ans = maxInt(ans, (r-l)*h)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
