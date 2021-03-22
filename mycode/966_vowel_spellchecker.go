package mycode

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	exactMatch := make(map[string]bool)
	caseMatch := make(map[string][]string)
	vowelMatch := make(map[string][]string)

	for _, w := range wordlist {
		wnc := strings.ToLower(w)
		wnv := spellcheckerReplaceVowel(wnc)

		exactMatch[w] = true

		if _, ok := caseMatch[wnc]; !ok {
			caseMatch[wnc] = make([]string, 0)
		}
		caseMatch[wnc] = append(caseMatch[wnc], w)

		if _, ok := vowelMatch[wnv]; !ok {
			vowelMatch[wnv] = make([]string, 0)
		}
		vowelMatch[wnv] = append(vowelMatch[wnv], w)
	}

	rst := make([]string, len(queries))
	for i, q := range queries {
		qnc := strings.ToLower(q)
		qnv := spellcheckerReplaceVowel(qnc)

		if _, ok := exactMatch[q]; ok {
			rst[i] = q
		} else if wl, ok := caseMatch[qnc]; ok {
			rst[i] = wl[0]
		} else if wl, ok := vowelMatch[qnv]; ok {
			rst[i] = wl[0]
		} else {
			rst[i] = ""
		}
	}

	return rst
}

func spellcheckerReplaceVowel(s string) string {
	buf := make([]rune, len(s))
	for i, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			buf[i] = '*'
		} else {
			buf[i] = c
		}
	}
	return string(buf)
}
