package mycode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	ans := ""
	pos := 0
	var ch uint8
out:
	for {
		for i, s := range strs {
			if pos >= len(s) {
				break out
			}
			if i == 0 {
				ch = strs[i][pos]
			} else if strs[i][pos] != ch {
				break out
			} else if i == len(strs)-1 {
				ans += string(rune(ch))
			} else {
				continue
			}
		}
		pos++
	}

	return ans
}
