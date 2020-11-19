package mycode

import "strings"

func decodeString(s string) string {
	numS := make([]int, 1, 30)
	sbS := make([]strings.Builder, 2, 30)

	// init two stacks, append ] to the end of s
	numS[0] = 1
	s += "]"
	ss := []rune(s)

	i := 0
	for i < len(s) && (len(numS) > 0 || len(sbS) > 1) {
		c := ss[i]

		if c >= 'a' && c <= 'z' {
			// alphabets
			sbS[len(sbS)-1].WriteRune(c)
			i++
		} else if c >= '0' && c <= '9' {
			// digits
			n := 0
			for i < len(s) && ss[i] >= '0' && ss[i] <= '9' {
				n = n*10 + int(ss[i]) - '0'
				i++
			}

			// after parsing digits, ss[i] should be '[', ignoring it
			i++

			// we're now inside a new [], appending new elements to stacks
			numS = append(numS, n)
			sbS = append(sbS, strings.Builder{})
		} else {
			// ']'
			i++

			// end of previous [], constructing result string
			prevN := numS[len(numS)-1]
			numS = numS[:len(numS)-1]
			prevSb := sbS[len(sbS)-1]
			sbS = sbS[:len(sbS)-1]
			for j := 0; j < prevN; j++ {
				sbS[len(sbS)-1].WriteString(prevSb.String())
			}
		}
	}

	// now len(sbS) should be 1 and sbS[0] should contain the final string
	return sbS[0].String()
}
