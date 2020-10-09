package mycode

type RecentCounter struct {
	buff [1e4]int
	l    int
	r    int
}

func Constructor933() RecentCounter {
	return RecentCounter{
		l: 0,
		r: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.buff[this.r] = t
	this.r++
	for this.buff[this.l] < t-3000 {
		this.l++
	}
	return this.r - this.l
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
