package mycode

// https://leetcode.com/problems/longest-string-chain/solution/
// Solution 1: DFS + memo
// Solution 2: Bottom-up DP
func longestStrChain(words []string) int {
	wordsMap := make(map[string]bool)
	memo := make(map[string]int)
	for _, w := range words {
		wordsMap[w] = true
	}

	ans := 0
	for _, w := range words {
		ans = maxInt(ans, longestStrChainDFS(wordsMap, memo, w))
	}
	return ans
}

func longestStrChainDFS(words map[string]bool, memo map[string]int, curWord string) int {
	if l, ok := memo[curWord]; ok {
		return l
	}

	maxLen := 1
	// work backwards, take out one char in curWord at a time
	for i := range curWord {
		newWord := curWord[:i] + curWord[i+1:]
		if _, ok := words[newWord]; ok {
			maxLen = maxInt(maxLen, 1+longestStrChainDFS(words, memo, newWord))
		}
	}
	memo[curWord] = maxLen
	return maxLen
}
