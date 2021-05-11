package mycode

// Solution 1: sliding window (convert to: sum-of-array - sum-of-window)
// Solution 2: DP, see below
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/discuss/598111/Java-dp-solution(explanation-with-picture)
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	m := n - k
	totSum := 0
	for _, t := range cardPoints {
		totSum += t
	}

	ans := 0
	curSum := 0
	for i, p := range cardPoints {
		curSum += p
		if i >= m {
			curSum -= cardPoints[i-m]
		}
		if i >= m-1 {
			ans = maxInt(ans, totSum-curSum)
		}
	}

	return ans
}
