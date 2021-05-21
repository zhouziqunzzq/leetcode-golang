package mycode

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	p := []rune(pattern)
	for _, w := range words {
		if findAndReplacePatternHelper([]rune(w), p) {
			ans = append(ans, w)
		}
	}
	return ans
}

func findAndReplacePatternHelper(word, pattern []rune) bool {
	if len(word) != len(pattern) {
		return false
	}

	w2p := make(map[rune]rune)
	p2w := make(map[rune]rune)

	for i := range word {
		w := word[i]
		p := pattern[i]
		if _, ok := w2p[w]; !ok {
			w2p[w] = p
		}
		if _, ok := p2w[p]; !ok {
			p2w[p] = w
		}
		if w2p[w] != p || p2w[p] != w {
			return false
		}
	}

	return true
}
