package mycode

func longestValidParentheses(s string) int {
	ans := 0
	stack := NewStack()
	stack.Push(-1)
	for i, v := range s {
		if v == '(' {
			stack.Push(i)
		} else if v == ')' {
			_ = stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				top, _ := stack.Peak().(int)
				if ans < i-top {
					ans = i - top
				}
			}
		}
	}
	return ans
}
