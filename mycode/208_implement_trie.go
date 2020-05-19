package mycode

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isWord bool
	edges  [26]*TrieNode
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.root
	wordArr := []rune(word)
	var i int

	for i = 0; i < len(wordArr); i++ {
		if p.edges[wordArr[i]-'a'] == nil {
			break
		} else {
			p = p.edges[wordArr[i]-'a']
		}
	}

	for ; i < len(wordArr); i++ {
		p.edges[wordArr[i]-'a'] = &TrieNode{}
		p = p.edges[wordArr[i]-'a']
	}
	p.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.root
	wordArr := []rune(word)

	for i := 0; i < len(wordArr); i++ {
		if p.edges[wordArr[i]-'a'] == nil {
			return false
		} else {
			p = p.edges[wordArr[i]-'a']
		}
	}
	return p.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	wordArr := []rune(prefix)

	for i := 0; i < len(wordArr); i++ {
		if p.edges[wordArr[i]-'a'] == nil {
			return false
		} else {
			p = p.edges[wordArr[i]-'a']
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
