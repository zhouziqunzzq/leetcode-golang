package mycode

func titleToNumber(s string) int {
	rst := 0
	for _, c := range []rune(s) {
		rst *= 26
		rst += int(c - 'A' + 1)
	}
	return rst
}
