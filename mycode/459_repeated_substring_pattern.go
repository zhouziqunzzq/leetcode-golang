package mycode

func repeatedSubstringPattern(s string) bool {
	for i := len(s) / 2; i >= 1; i-- {
		if len(s)%i == 0 {
			subStr := s[:i]
			k := 1
			m := len(s) / i
			for ; k < m; k++ {
				if subStr != s[k*i:(k+1)*i] {
					break
				}
			}
			if k == m {
				return true
			}
		}
	}
	return false
}
