package mycode

func halvesAreAlike(s string) bool {
	n := len(s)
	return halvesAreAlikeHelper(s[:n/2]) == halvesAreAlikeHelper(s[n/2:])
}

func halvesAreAlikeHelper(s string) int {
	vowels := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			vowels++
		}
	}
	return vowels
}
