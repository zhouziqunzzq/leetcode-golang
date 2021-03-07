package mycode

type HashMapListNode struct {
	K    int
	V    int
	Next *HashMapListNode
}

type MyHashMap struct {
	buckets [9001]*HashMapListNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Find(key int) (prev, p *HashMapListNode) {
	prev = nil
	p = this.buckets[key%9001]
	for p != nil {
		if p.K == key {
			return
		}
		prev = p
		p = p.Next
	}
	return
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	_, p := this.Find(key)
	if p != nil {
		p.V = value
	} else {
		idx := key % 9001
		this.buckets[idx] = &HashMapListNode{key, value, this.buckets[idx]}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	_, p := this.Find(key)
	if p != nil {
		return p.V
	} else {
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	idx := key % 9001
	prev, p := this.Find(key)
	if p == nil {
		return
	}
	if prev == nil {
		this.buckets[idx] = p.Next
	} else {
		prev.Next = p.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
