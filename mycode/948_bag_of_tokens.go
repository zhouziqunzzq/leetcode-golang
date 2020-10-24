package mycode

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})

	maxScore := 0
	curScore := 0
	l, r := 0, len(tokens)-1

	for l <= r {
		if P >= tokens[l] {
			P -= tokens[l]
			curScore += 1
			l++
			maxScore = maxInt(maxScore, curScore)
		} else if P < tokens[l] && curScore > 0 {
			curScore--
			P += tokens[r]
			r--
		} else {
			return maxScore
		}
	}
	return maxScore
}
