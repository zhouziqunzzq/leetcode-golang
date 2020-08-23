package mycode

import "strings"

type StreamChecker struct {
	root *TrieNode
	buf  strings.Builder
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{root: &TrieNode{}}
	for _, s := range words {
		node := sc.root
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i]
			if node.edges[c-'a'] == nil {
				node.edges[c-'a'] = &TrieNode{}
			}
			node = node.edges[c-'a']
		}
		node.isWord = true
	}
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.buf.WriteByte(letter)
	node := this.root
	s := this.buf.String()
	for i := len(s) - 1; i >= 0 && node != nil; i-- {
		c := s[i]
		node = node.edges[c-'a']
		if node != nil && node.isWord {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
