package mycode

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	mSet  map[int]int
	mNums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		mSet:  make(map[int]int),
		mNums: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mSet[val]; ok {
		return false
	} else {
		this.mNums = append(this.mNums, val)
		this.mSet[val] = len(this.mNums) - 1
		return true
	}
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.mSet[val]; !ok {
		return false
	} else {
		this.mNums[idx] = this.mNums[len(this.mNums)-1]
		this.mSet[this.mNums[len(this.mNums)-1]] = idx
		this.mNums = this.mNums[:len(this.mNums)-1]
		delete(this.mSet, val)
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.mNums[rand.Intn(len(this.mNums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
