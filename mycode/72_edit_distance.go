package mycode

func minDistance72(word1 string, word2 string) int {
	opt := make([][]int, len(word1)+1)
	for i, _ := range opt {
		opt[i] = make([]int, len(word2)+1)
	}

	// base case
	for i := 0; i <= len(word1); i++ {
		opt[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		opt[0][j] = j
	}

	// fill the table
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i] == word2[j] {
				opt[i][j] = opt[i-1][j-1]
			} else {
				opt[i][j] = minInt(minInt(opt[i-1][j], opt[i][j-1]), opt[i-1][j-1]) + 1
			}
		}
	}

	return opt[len(word1)][len(word2)]
}
