package mycode

// https://leetcode.com/problems/iterator-for-combination/discuss/451544/Java-No-pre-calculation-needed-for-iterator-questions
type CombinationIterator struct {
	ch2Idx  map[rune]int
	rst     []rune
	str     []rune
	combLen int
	rstIdx  int
}

func Constructor1286(characters string, combinationLength int) CombinationIterator {
	ci := CombinationIterator{
		ch2Idx:  make(map[rune]int),
		rst:     make([]rune, combinationLength),
		str:     []rune(characters),
		combLen: combinationLength,
		rstIdx:  combinationLength - 1,
	}
	for i, c := range ci.str {
		ci.ch2Idx[c] = i
		if i < combinationLength {
			ci.rst[i] = c
		}
	}
	return ci
}

func (this *CombinationIterator) Next() string {
	curRst := make([]rune, this.combLen)
	for i, c := range this.rst {
		curRst[i] = c
	}

	idx := len(this.str) - 1
	this.rstIdx = len(this.rst) - 1
	for this.rstIdx >= 0 && this.ch2Idx[this.rst[this.rstIdx]] == idx {
		idx--
		this.rstIdx--
	}
	if this.rstIdx < 0 { // no next result to pre-process
		return string(curRst)
	}

	idx = this.ch2Idx[this.rst[this.rstIdx]]
	rstIdx := this.rstIdx
	for rstIdx < len(this.rst) {
		idx++
		this.rst[rstIdx] = this.str[idx]
		rstIdx++
	}

	return string(curRst)
}

func (this *CombinationIterator) HasNext() bool {
	return this.rstIdx >= 0
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
