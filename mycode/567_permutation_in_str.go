package mycode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var key1, key2 [26]int

	for _, cp := range s1 {
		key1[cp-'a']++
	}
	for i := 0; i < len(s1); i++ {
		key2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		isAnagram := true
		for j := 0; j < 26; j++ {
			if key1[j] != key2[j] {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			return true
		}

		if i < len(s2)-len(s1) {
			key2[s2[i]-'a']--
			key2[s2[i+len(s1)]-'a']++
		}
	}
	return false
}
