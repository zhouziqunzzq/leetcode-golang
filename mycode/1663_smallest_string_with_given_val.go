package mycode

// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/discuss/944411/C%2B%2B-Simple-O(N)-with-explanation
// Solution: greedily build string from right to left
func getSmallestString(n int, k int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = 'a'
	}
	k -= n

	cur := n - 1
	for k > 0 {
		gain := minInt(k, 25)
		s[cur] += rune(gain)
		cur--
		k -= gain
	}

	return string(s)
}
