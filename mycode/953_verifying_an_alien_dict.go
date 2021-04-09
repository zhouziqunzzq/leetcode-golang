package mycode

func isAlienSorted(words []string, order string) bool {
	c2idx := make(map[rune]int)
	for i, c := range []rune(order) {
		c2idx[c] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !isAlienSortedCmp(words[i], words[i+1], c2idx) {
			return false
		}
	}
	return true
}

// return true iff. a <= b
func isAlienSortedCmp(a, b string, c2idx map[rune]int) bool {
	m, n := len(a), len(b)
	i, j := 0, 0
	for i < m && j < n {
		if c2idx[rune(a[i])] < c2idx[rune(b[j])] {
			return true
		} else if c2idx[rune(a[i])] > c2idx[rune(b[j])] {
			return false
		}
		i++
		j++
	}
	return m <= n
}
