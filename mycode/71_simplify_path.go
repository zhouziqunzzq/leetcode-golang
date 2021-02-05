package mycode

import "strings"

func simplifyPath(path string) string {
	s := make([]string, 0)
	sb := strings.Builder{}
	n := len(path)

	for i, c := range []rune(path) {
		if c != '/' {
			sb.WriteRune(c)
		}
		if c == '/' || i == n-1 {
			if sb.Len() > 0 {
				buf := sb.String()
				sb.Reset()
				switch buf {
				case ".":
					continue
				case "..":
					if len(s) > 0 {
						s = s[:len(s)-1]
					}
				default:
					s = append(s, buf)
				}
			}
		}
	}

	if len(s) == 0 {
		return "/"
	} else {
		sb.Reset()
		for _, p := range s {
			sb.WriteRune('/')
			sb.WriteString(p)
		}
		return sb.String()
	}
}
