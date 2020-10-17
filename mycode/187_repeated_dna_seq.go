package mycode

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	ans := make([]string, 0)
	for i := 0; i+9 < len(s); i++ {
		pat := s[i : i+10]
		if n, ok := m[pat]; ok {
			if n == 0 {
				ans = append(ans, pat)
			}
			m[pat]++
		} else {
			m[pat] = 0
		}
	}

	return ans
}
