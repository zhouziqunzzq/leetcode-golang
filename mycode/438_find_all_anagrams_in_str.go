package mycode

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)

	if len(s) < len(p) {
		return ans
	}

	var keyS, keyP [26]int

	for _, cp := range p {
		keyP[cp-'a']++
	}
	for i := 0; i < len(p); i++ {
		keyS[s[i]-'a']++
	}

	for i := 0; i < len(s)-len(p)+1; i++ {
		isAnagram := true
		for j := 0; j < 26; j++ {
			if keyS[j] != keyP[j] {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			ans = append(ans, i)
		}

		if i < len(s)-len(p) {
			keyS[s[i]-'a']--
			keyS[s[i+len(p)]-'a']++
		}
	}
	return ans
}
