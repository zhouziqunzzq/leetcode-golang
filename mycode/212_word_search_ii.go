package mycode

type WordTrieNode struct {
	next [26]*WordTrieNode
	word string
}

func buildWordTrieNode(words []string) (root *WordTrieNode) {
	root = &WordTrieNode{}
	for _, w := range words {
		p := root
		for _, c := range w {
			i := c - 'a'
			if p.next[i] == nil {
				p.next[i] = &WordTrieNode{}
			}
			p = p.next[i]
		}
		p.word = w
	}
	return
}

func findWords(board [][]byte, words []string) []string {
	rst := make([]string, 0)
	root := buildWordTrieNode(words)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			findWordsHelper(board, i, j, root, &rst)
		}
	}

	return rst
}

func findWordsHelper(board [][]byte, i, j int, p *WordTrieNode, rst *[]string) {
	c := board[i][j]

	if c == '#' || p.next[c-'a'] == nil { // early stopping (pruning)
		return
	}

	p = p.next[c-'a']
	if len(p.word) > 0 { // found one word
		*rst = append(*rst, p.word)
		p.word = "" // prevent duplicate
	}

	board[i][j] = '#' // visit flag
	if i > 0 {
		findWordsHelper(board, i-1, j, p, rst)
	}
	if i < len(board)-1 {
		findWordsHelper(board, i+1, j, p, rst)
	}
	if j > 0 {
		findWordsHelper(board, i, j-1, p, rst)
	}
	if j < len(board[0])-1 {
		findWordsHelper(board, i, j+1, p, rst)
	}
	board[i][j] = c // back-tracing
}
