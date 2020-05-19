package mycode

type StockSpannerElem struct {
	Price int
	Span  int
}
type StockSpanner struct {
	mStack []StockSpannerElem // monotonous stack
}

func Constructor() StockSpanner {
	return StockSpanner{
		mStack: make([]StockSpannerElem, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	ans := 1

	for len(this.mStack) > 0 && price >= this.mStack[len(this.mStack)-1].Price {
		ans += this.mStack[len(this.mStack)-1].Span
		this.mStack = this.mStack[:len(this.mStack)-1]
	}

	this.mStack = append(this.mStack, StockSpannerElem{price, ans})

	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
