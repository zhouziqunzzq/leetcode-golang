package mycode

import "strings"

func uniqueMorseRepresentations(words []string) int {
	t := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-",
		".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--.."}
	ans := make(map[string]bool)
	sb := strings.Builder{}

	for _, w := range words {
		sb.Reset()
		for _, c := range []rune(w) {
			sb.WriteString(t[c-'a'])
		}
		s := sb.String()
		if _, ok := ans[s]; !ok {
			ans[s] = true
		}
	}

	return len(ans)
}
