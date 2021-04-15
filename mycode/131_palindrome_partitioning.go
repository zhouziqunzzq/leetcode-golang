package mycode

func partition131(s string) [][]string {
	buf := make([]string, 0)
	rst := make([][]string, 0)
	partitionHelper(s, 0, &buf, &rst)
	return rst
}

func partitionHelper(s string, l int, buf *[]string, rst *[][]string) {
	// base case
	if l >= len(s) {
		t := make([]string, len(*buf))
		for i := range t {
			t[i] = (*buf)[i]
		}
		*rst = append(*rst, t)
	}

	// dfs
	for r := l; r < len(s); r++ {
		if isPalindrome131(s[l : r+1]) {
			*buf = append(*buf, s[l:r+1])
			partitionHelper(s, r+1, buf, rst)
			*buf = (*buf)[:len(*buf)-1]
		}
	}
}

// can be further optimized with DP to omit overlapping palindrome checks
func isPalindrome131(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
