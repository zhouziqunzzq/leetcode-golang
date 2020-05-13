package mycode

// monotonous stack
func removeKdigits(num string, k int) string {
	ans := make([]rune, 0)

	for _, c := range num {
		for k > 0 && len(ans) > 0 && ans[len(ans)-1] > c {
			ans = ans[:len(ans)-1]
			k--
		}
		if len(ans) > 0 || c != '0' {
			ans = append(ans, c)
		}
	}

	if k > 0 {
		if k <= len(ans) {
			ans = ans[:len(ans)-k]
		} else {
			return "0"
		}
	}
	if len(ans) > 0 {
		return string(ans)
	} else {
		return "0"
	}
}
