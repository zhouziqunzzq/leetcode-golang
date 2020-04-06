package mycode

func isValid_1(s string) bool {
	st := NewStack()
	for _, ss := range s {
		switch ss {
		case '(':
			st.Push(')')
		case '{':
			st.Push('}')
		case '[':
			st.Push(']')
		case ')':
			fallthrough
		case '}':
			fallthrough
		case ']':
			if st.Empty() {
				return false
			}
			top := st.Pop().(rune)
			if top != ss {
				return false
			}
		}
	}

	return st.Empty()
}
