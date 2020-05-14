package mycode

import "container/list"

type FirstUnique struct {
	uniqueList *list.List
	elemMap    map[int]*list.Element
}

func FirstUniqueConstructor(nums []int) FirstUnique {
	fu := FirstUnique{uniqueList: list.New(), elemMap: make(map[int]*list.Element)}

	for _, v := range nums {
		if elemPtr, ok := fu.elemMap[v]; ok {
			if elemPtr != nil {
				fu.uniqueList.Remove(elemPtr)
				fu.elemMap[v] = nil
			}
		} else {
			fu.elemMap[v] = fu.uniqueList.PushBack(v)
		}
	}

	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.uniqueList.Len() == 0 {
		return -1
	} else {
		return this.uniqueList.Front().Value.(int)
	}
}

func (this *FirstUnique) Add(value int) {
	if elemPtr, ok := this.elemMap[value]; ok {
		if elemPtr != nil {
			this.uniqueList.Remove(elemPtr)
			this.elemMap[value] = nil
		}
	} else {
		this.elemMap[value] = this.uniqueList.PushBack(value)
	}
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
