package mycode

// https://leetcode.com/problems/prefix-and-suffix-search/solution/
// Solution 3: Trie of suffix wrapped words.
// e.g. consider the word 'apple', we'll insert '#apple',
// 'e#apple', 'le#apple', 'ple#apple', 'pple#apple',
// 'apple#apple' into the trie.

type TrieNode745 struct {
	children [27]*TrieNode745
	weight   int
}

type WordFilter struct {
	trie *TrieNode745
}

func Constructor745(words []string) WordFilter {
	wf := WordFilter{&TrieNode745{}}
	trie := wf.trie

	for weight := 0; weight < len(words); weight++ {
		// '{' is the very next char after 'z'
		word := words[weight] + "{"
		for i := 0; i < len(word); i++ {
			cur := trie
			cur.weight = weight
			for j := i; j < 2*len(word)-1; j++ {
				k := int(word[j%len(word)] - 'a')
				if cur.children[k] == nil {
					cur.children[k] = &TrieNode745{}
				}
				cur = cur.children[k]
				cur.weight = weight
			}
		}
	}

	return wf
}

func (this *WordFilter) F(prefix string, suffix string) int {
	cur := this.trie
	for _, c := range suffix + "{" + prefix {
		if cur.children[int(c-'a')] == nil {
			return -1
		}
		cur = cur.children[int(c-'a')]
	}
	return cur.weight
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
