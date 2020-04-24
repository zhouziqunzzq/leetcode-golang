package mycode

import "container/list"

type LRUCacheEntry struct {
	Val  int
	Elem *list.Element
}

type LRUCache struct {
	cache    map[int]LRUCacheEntry
	keyList  *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		make(map[int]LRUCacheEntry),
		list.New(),
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.touch(key)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if oldV, ok := this.cache[key]; ok {
		this.touch(key)
		this.cache[key] = LRUCacheEntry{value, oldV.Elem}
	} else {
		if len(this.cache) == this.capacity {
			oldKey := this.keyList.Remove(this.keyList.Back()).(int)
			delete(this.cache, oldKey)
		}
		this.cache[key] = LRUCacheEntry{value, this.keyList.PushFront(key)}
	}
}

func (this *LRUCache) touch(key int) {
	this.keyList.MoveToFront(this.cache[key].Elem)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
