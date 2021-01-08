package mycode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	p1, pp1 := 0, 0
	p2, pp2 := 0, 0

	for {
		if word1[p1][pp1] != word2[p2][pp2] {
			return false
		}

		pp1++
		if pp1 == len(word1[p1]) {
			p1++
			pp1 = 0
		}

		pp2++
		if pp2 == len(word2[p2]) {
			p2++
			pp2 = 0
		}

		if p1 == len(word1) || p2 == len(word2) {
			break
		}
	}

	return p1 == len(word1) && p2 == len(word2) && pp1 == 0 && pp2 == 0
}
