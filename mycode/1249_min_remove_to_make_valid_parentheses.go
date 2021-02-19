package mycode

import "strings"

func minRemoveToMakeValid(s string) string {
	ss := []rune(s)

	// forward pass
	openCnt := 0
	for i, c := range ss {
		if c == '(' {
			openCnt++
		} else if c == ')' {
			openCnt--
		} else {
			continue
		}
		if openCnt < 0 {
			ss[i] = '-'
			openCnt = 0
		}
	}

	// backward pass
	openCnt = 0
	for i := len(ss) - 1; i >= 0; i-- {
		c := ss[i]
		if c == ')' {
			openCnt++
		} else if c == '(' {
			openCnt--
		} else {
			continue
		}
		if openCnt < 0 {
			ss[i] = '-'
			openCnt = 0
		}
	}

	sb := strings.Builder{}
	for _, c := range ss {
		if c != '-' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
