package mycode

import "strconv"

func evalRPN(tokens []string) int {
	s := make([]int, 0)

	for _, tkn := range tokens {
		if len(tkn) > 1 || tkn[0] >= '0' && tkn[0] <= '9' {
			v, _ := strconv.Atoi(tkn)
			s = append(s, v)
		} else {
			a, b := s[len(s)-2], s[len(s)-1]
			s = s[:len(s)-2]
			t := 0
			switch tkn {
			case "+":
				t = a + b
			case "-":
				t = a - b
			case "*":
				t = a * b
			case "/":
				t = a / b
			}
			s = append(s, t)
		}
	}

	return s[len(s)-1]
}
