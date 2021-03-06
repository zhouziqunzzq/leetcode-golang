package mycode

type TrieNode820 struct {
	Cnt   int
	Edges [26]*TrieNode820
}

func minimumLengthEncoding(words []string) int {
	root := &TrieNode820{}
	nodeToIdx := make(map[*TrieNode820]int)

	for idx, w := range words {
		cur := root
		for i := len(w) - 1; i >= 0; i-- {
			c := w[i]
			cur.Cnt++
			if cur.Edges[c-'a'] == nil {
				cur.Edges[c-'a'] = &TrieNode820{}
			}
			cur = cur.Edges[c-'a']
		}
		nodeToIdx[cur] = idx
	}

	ans := 0
	for node, idx := range nodeToIdx {
		if node.Cnt == 0 {
			ans += len(words[idx]) + 1
		}
	}

	return ans
}
