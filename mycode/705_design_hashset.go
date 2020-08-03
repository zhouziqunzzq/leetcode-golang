package mycode

type MyHashSet struct {
	buckets [9001]*ListNode
}

/** Initialize your data structure here. */
func HashSetConstructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

	idx := key % 9001
	this.buckets[idx] = &ListNode{key, this.buckets[idx]}
}

func (this *MyHashSet) Find(key int) (prev, p *ListNode) {
	prev = nil
	p = this.buckets[key%9001]
	for p != nil {
		if p.Val == key {
			return
		}
		prev = p
		p = p.Next
	}
	return
}

func (this *MyHashSet) Remove(key int) {
	idx := key % 9001
	prev, p := this.Find(key)
	if p == nil {
		return
	}
	if prev == nil { // remove head
		this.buckets[idx] = p.Next
	} else { // remove inner node
		prev.Next = p.Next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, p := this.Find(key)
	return p != nil
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
