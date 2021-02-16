package mycode

import (
	"unicode"
)

func letterCasePermutation(S string) []string {
	rst := make([]string, 0)
	buf := make([]rune, len(S))
	letterCasePermutationHelper([]rune(S), buf, 0, &rst)
	return rst
}

func letterCasePermutationHelper(s []rune, buf []rune, curPos int, rst *[]string) {
	if curPos >= len(s) {
		t := make([]rune, len(buf))
		for i := range buf {
			t[i] = buf[i]
		}
		*rst = append(*rst, string(t))
		return
	}

	if s[curPos] >= '0' && s[curPos] <= '9' {
		buf[curPos] = s[curPos]
		letterCasePermutationHelper(s, buf, curPos+1, rst)
	} else {
		lc := unicode.ToLower(s[curPos])
		uc := unicode.ToUpper(s[curPos])
		for _, c := range []rune{lc, uc} {
			buf[curPos] = c
			letterCasePermutationHelper(s, buf, curPos+1, rst)
		}
	}
}
