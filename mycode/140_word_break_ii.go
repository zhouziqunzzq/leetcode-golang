package mycode

import "strings"

func wordBreak2(s string, wordDict []string) []string {
	return wordBreakHelper(s, wordDict, make(map[string][]string))
}

// memorized DFS
func wordBreakHelper(s string, wordDict []string, mem map[string][]string) []string {
	// check memory
	if rst, ok := mem[s]; ok {
		return rst
	}

	rst := make([]string, 0)

	// base case
	if len(s) == 0 {
		return append(rst, "")
	}

	// dfs
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			subRst := wordBreakHelper(s[len(word):], wordDict, mem)
			for _, sub := range subRst {
				if len(sub) > 0 {
					rst = append(rst, word+" "+sub)
				} else {
					rst = append(rst, word)
				}
			}
		}
	}

	// update memory
	mem[s] = rst

	return rst
}
