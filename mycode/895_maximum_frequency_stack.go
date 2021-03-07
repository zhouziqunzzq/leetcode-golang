package mycode

// https://leetcode.com/problems/maximum-frequency-stack/
type FreqStack struct {
	elem2Freq      map[int]int
	freq2ElemStack map[int][]int
	maxFreq        int
}

func Constructor895() FreqStack {
	return FreqStack{
		elem2Freq:      make(map[int]int),
		freq2ElemStack: make(map[int][]int),
		maxFreq:        0,
	}
}

func (this *FreqStack) Push(x int) {
	// increase freq count
	if _, ok := this.elem2Freq[x]; !ok {
		this.elem2Freq[x] = 0
	}
	this.elem2Freq[x]++
	f := this.elem2Freq[x]

	// push to elem stack of that freq
	if _, ok := this.freq2ElemStack[f]; !ok {
		this.freq2ElemStack[f] = make([]int, 0)
	}
	this.freq2ElemStack[f] = append(this.freq2ElemStack[f], x)

	// update max freq
	this.maxFreq = maxInt(this.maxFreq, f)
}

func (this *FreqStack) Pop() int {
	// get elem from elem stack of max freq
	maxFreqStack := this.freq2ElemStack[this.maxFreq]
	x := maxFreqStack[len(maxFreqStack)-1]
	this.freq2ElemStack[this.maxFreq] = this.freq2ElemStack[this.maxFreq][:len(maxFreqStack)-1]

	// update elem to freq
	this.elem2Freq[x]--

	// update max freq
	if len(this.freq2ElemStack[this.maxFreq]) == 0 {
		this.maxFreq--
	}

	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
