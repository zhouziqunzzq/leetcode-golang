package mycode

func calculate(s string) int {
	p := make(map[rune]uint) // p[i]: priority of operator i
	p['#'] = 0
	p['+'] = 1
	p['-'] = 1
	p['*'] = 2
	p['/'] = 2

	s += "#"
	ss := []rune(s)

	ops := make([]rune, 0)
	ns := make([]int, 0)

	for i := 0; i < len(ss); {
		if ss[i] == ' ' {
			i++
		} else if ss[i] >= '0' && ss[i] <= '9' {
			t := 0
			for i < len(ss) && ss[i] >= '0' && ss[i] <= '9' {
				t = t*10 + int(ss[i]-'0')
				i++
			}
			ns = append(ns, t)
		} else {
			curOp := ss[i]
			i++

			for len(ops) > 0 && p[ops[len(ops)-1]] >= p[curOp] {
				topOp := ops[len(ops)-1]
				b := ns[len(ns)-1]
				a := ns[len(ns)-2]
				ns = ns[:len(ns)-2]

				rst := 0
				if topOp == '+' {
					rst = a + b
				} else if topOp == '-' {
					rst = a - b
				} else if topOp == '*' {
					rst = a * b
				} else {
					rst = a / b
				}
				ns = append(ns, rst)

				ops = ops[:len(ops)-1]
			}

			ops = append(ops, curOp)
		}
	}

	if len(ns) > 0 {
		return ns[0]
	} else {
		return 0
	}
}
