package mycode

import "strings"

func removeDuplicates(s string, k int) string {
	ss := make([]rune, 0)
	sc := make([]int, 0)

	for _, c := range []rune(s) {
		if len(ss) == 0 || ss[len(ss)-1] != c {
			ss = append(ss, c)
			sc = append(sc, 1)
		} else {
			sc[len(sc)-1]++
			if sc[len(sc)-1] >= k {
				ss = ss[:len(ss)-1]
				sc = sc[:len(sc)-1]
			}
		}
	}

	sb := strings.Builder{}
	for i := range ss {
		for j := 0; j < sc[i]; j++ {
			sb.WriteRune(ss[i])
		}
	}

	return sb.String()
}
