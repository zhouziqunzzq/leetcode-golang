package mycode

import "strings"

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	// pattern -> string
	f1 := make(map[rune]string)
	// string -> pattern
	f2 := make(map[string]rune)
	for i, c := range []rune(pattern) {
		s, ok1 := f1[c]
		p, ok2 := f2[strs[i]]
		if !ok1 && !ok2 {
			f1[c] = strs[i]
			f2[strs[i]] = c
		} else if ok1 && ok2 && s == strs[i] && p == c {
			continue
		} else {
			return false
		}
	}
	return true
}
