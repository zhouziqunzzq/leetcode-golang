package mycode

func findLongestWord(s string, d []string) string {
	rst := ""

	for _, ds := range d {
		if findLongestWordHelper(ds, s) {
			if len(ds) > len(rst) || (len(ds) == len(rst) && ds < rst) {
				rst = ds
			}
		}
	}

	return rst
}

// check whether s1 is a subsequence of s2
func findLongestWordHelper(s1, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}
