package mycode

func lengthOfLongestSubstring(s string) int {
	str := []rune(s)
	posMap := make(map[rune]int)
	ans := 0
	for i, j := 0, 0; j < len(str); j++ {
		if pos, ok := posMap[str[j]]; ok {
			if pos > i {
				i = pos
			}
		}
		t := j - i + 1
		if t > ans {
			ans = t
		}
		posMap[str[j]] = j + 1
	}
	return ans
}
