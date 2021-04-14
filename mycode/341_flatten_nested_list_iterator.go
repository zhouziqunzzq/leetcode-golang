package mycode

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool           { return true }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	lStack [][]*NestedInteger
	pStack []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	t := &NestedIterator{make([][]*NestedInteger, 1), make([]int, 1)}
	t.lStack[0] = nestedList
	t.pStack[0] = -1
	t.prepareNext()
	return t
}

func (this *NestedIterator) moveAndPop() {
	// move pointers and pop out exhausted lists
	for len(this.lStack) > 0 && len(this.pStack) > 0 {
		this.pStack[len(this.pStack)-1]++
		if this.pStack[len(this.pStack)-1] >= len(this.lStack[len(this.lStack)-1]) {
			this.lStack = this.lStack[:len(this.lStack)-1]
			this.pStack = this.pStack[:len(this.pStack)-1]
		} else {
			break
		}
	}
}

func (this *NestedIterator) prepareNext() {
	this.moveAndPop()

	for len(this.lStack) > 0 && len(this.pStack) > 0 {
		curNext := this.lStack[len(this.lStack)-1][this.pStack[len(this.pStack)-1]]
		if curNext.IsInteger() {
			break
		} else {
			nextList := curNext.GetList()
			if len(nextList) > 0 {
				this.lStack = append(this.lStack, nextList)
				this.pStack = append(this.pStack, 0)
			} else {
				this.moveAndPop()
			}
		}
	}
}

func (this *NestedIterator) Next() int {
	if !this.HasNext() {
		return -1
	}
	t := this.lStack[len(this.lStack)-1][this.pStack[len(this.pStack)-1]].GetInteger()
	this.prepareNext()
	return t
}

func (this *NestedIterator) HasNext() bool {
	return len(this.lStack) > 0 && len(this.pStack) > 0
}
