package mycode

// pseudo func
func rand7() int {
	return 0
}

// https://leetcode.com/problems/implement-rand10-using-rand7/solution/
func rand10() int {
	idx := 41
	for idx > 40 {
		row, col := rand7(), rand7()
		idx = row + (col-1)*7
	}
	return (idx-1)%10 + 1
}
