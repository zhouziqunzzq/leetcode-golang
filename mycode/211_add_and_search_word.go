package mycode

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor211() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	n := this.root
	for _, c := range []rune(word) {
		if n.edges[c-'a'] == nil {
			n.edges[c-'a'] = &TrieNode{}
		}
		n = n.edges[c-'a']
	}
	n.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match([]rune(word), 0, this.root)
}

func (this *WordDictionary) match(word []rune, pos int, n *TrieNode) bool {
	if pos >= len(word) {
		return n.isWord
	}
	if word[pos] != '.' {
		return n.edges[word[pos]-'a'] != nil && this.match(word, pos+1, n.edges[word[pos]-'a'])
	} else {
		for _, e := range n.edges {
			if e != nil && this.match(word, pos+1, e) {
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
