package mycode

type MyCircularQueue struct {
	elem     []int
	size     int
	capacity int
	fp       int // front pointer
	rp       int // rear pointer
}

func Constructor622(k int) MyCircularQueue {
	return MyCircularQueue{
		elem:     make([]int, k),
		size:     0,
		capacity: k,
		fp:       0,
		rp:       0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.elem[this.rp] = value
	this.size++
	this.rp = (this.rp + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.size--
	this.fp = (this.fp + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.elem[this.fp]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		trp := this.rp - 1
		if trp < 0 {
			trp += this.capacity
		}
		return this.elem[trp]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
