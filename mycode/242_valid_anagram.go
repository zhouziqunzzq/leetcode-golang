package mycode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var cnt1, cnt2 [26]int
	for i, c := range s {
		cnt1[c-'a']++
		cnt2[t[i]-'a']++
	}

	for i := range cnt1 {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}
