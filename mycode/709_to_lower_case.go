package mycode

func toLowerCase(s string) string {
	ss := []rune(s)
	for i, c := range ss {
		if c >= 'A' && c <= 'Z' {
			ss[i] = c + 'a' - 'A'
		}
	}
	return string(ss)
}
