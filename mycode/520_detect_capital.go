package mycode

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}

	allUpper := true
	allLower := true
	firstUpper := word[0] >= 'A' && word[0] <= 'Z'

	for _, c := range word[1:] {
		allUpper = allUpper && (c >= 'A' && c <= 'Z')
		allLower = allLower && (c >= 'a' && c <= 'z')
	}

	return (firstUpper && allUpper) || allLower
}
