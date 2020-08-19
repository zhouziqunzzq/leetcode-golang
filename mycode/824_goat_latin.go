package mycode

import "strings"

func toGoatLatin(S string) string {
	var sb strings.Builder
	ss := strings.Split(S, " ")

	for i, s := range ss {
		beg := s[0]
		if beg == 'a' || beg == 'e' || beg == 'i' || beg == 'o' || beg == 'u' ||
			beg == 'A' || beg == 'E' || beg == 'I' || beg == 'O' || beg == 'U' {
			sb.WriteString(s)
		} else {
			if len(s) > 1 {
				sb.WriteString(s[1:])
			}
			sb.WriteRune(rune(beg))
		}

		sb.WriteString("ma")
		for j := 0; j < i+1; j++ {
			sb.WriteRune('a')
		}

		if i < len(ss)-1 {
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}
