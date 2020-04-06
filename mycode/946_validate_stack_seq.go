package mycode

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	const (
		Invalid = 0
		Valid   = 1
		Pushed  = 2
	)

	if len(pushed) != len(popped) {
		return false
	}
	pushedFlags := make([]int8, 1000)
	for i := range pushedFlags {
		pushedFlags[i] = Invalid
	}
	for _, v := range pushed {
		pushedFlags[v] = Valid
	}

	s := NewStack()
	pushedIndex := 0

	for _, v := range popped {
		if pushedFlags[v] == Invalid {
			return false
		} else if pushedFlags[v] == Valid {
			for pushed[pushedIndex] != v {
				p := pushed[pushedIndex]
				s.Push(p)
				pushedFlags[p] = Pushed
				pushedIndex++
			}
			pushedFlags[pushed[pushedIndex]] = Pushed
			pushedIndex++
		} else if pushedFlags[v] == Pushed {
			found := false
			for !s.Empty() {
				pp := s.Pop().(int)
				if pp == v {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
