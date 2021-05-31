package mycode

// Solution 1: sort + bin-search
// Solution 2: sort + 2 ptrs
// Solution 3: Trie + DFS  ***
func suggestedProducts(products []string, searchWord string) [][]string {
	trie := TrieConstructor()
	for _, p := range products {
		trie.Insert(p)
	}

	ans := make([][]string, 0)
	for i := 1; i <= len(searchWord); i++ {
		rstBuf := make([]string, 0)
		node := trie.WalkDown(searchWord[:i])
		if node != nil {
			trie.SuggestedProductsDFS(node, searchWord[:i], &rstBuf)
		}
		ans = append(ans, rstBuf)
	}
	return ans
}

func (t *Trie) WalkDown(s string) *TrieNode {
	p := t.root
	for _, c := range s {
		if p.edges[c-'a'] == nil {
			return nil
		}
		p = p.edges[c-'a']
	}
	return p
}

func (t *Trie) SuggestedProductsDFS(curr *TrieNode, buf string, rst *[]string) {
	if len(*rst) >= 3 {
		return
	}

	if curr.isWord {
		*rst = append(*rst, buf)
	}
	for c := 'a'; c <= 'z'; c++ {
		if curr.edges[c-'a'] != nil {
			t.SuggestedProductsDFS(curr.edges[c-'a'], buf+string(c), rst)
		}
	}
}
