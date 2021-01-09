package mycode

// https://leetcode.com/problems/word-ladder/solution/
// Solution 1: BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	pToIds := make(map[string][]int) // generic pattern to list of IDs of word in wordList
	for id, w := range wordList {
		for i := range w {
			p := w[:i] + "*" + w[i+1:]
			if _, ok := pToIds[p]; !ok {
				pToIds[p] = make([]int, 0)
			}
			pToIds[p] = append(pToIds[p], id)
		}
	}

	wordList = append(wordList, beginWord)
	defer func() {
		wordList = wordList[:len(wordList)-1]
	}()

	visited := make(map[int]bool)
	q := make([]int, 1)
	q[0] = len(wordList) - 1
	d := 0
	for len(q) > 0 {
		d++
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			w := wordList[q[i]]
			// check ending condition
			if w == endWord {
				return d
			}
			// expand this node
			for j := 0; j < len(w); j++ {
				p := w[:j] + "*" + w[j+1:]
				if nextIdList, ok := pToIds[p]; ok {
					for _, nextId := range nextIdList {
						if _, ok := visited[nextId]; !ok {
							q = append(q, nextId)
							visited[nextId] = true
						}
					}
				}
			}
		}
		q = q[curLen:]
	}

	return 0
}
