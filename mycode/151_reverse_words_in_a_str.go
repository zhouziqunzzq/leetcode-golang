package mycode

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(strings.TrimSpace(s), " ")
	var buf bytes.Buffer
	for i := len(ss) - 1; i >= 0; i-- {
		if len(ss[i]) > 0 {
			buf.WriteString(ss[i])
			if i > 0 {
				buf.WriteString(" ")
			}
		}
	}
	return buf.String()
}
