package mycode

// can be further improved by using []int and custom hash function instead of using map...
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 && len(magazine) == 0 {
		return true
	} else if len(ransomNote) > len(magazine) {
		return false
	}

	letterMap := make(map[rune]int)

	for _, l := range magazine {
		if _, ok := letterMap[l]; ok {
			letterMap[l]++
		} else {
			letterMap[l] = 1
		}
	}

	for _, l := range ransomNote {
		if v, ok := letterMap[l]; v > 0 && ok {
			letterMap[l]--
		} else {
			return false
		}
	}

	return true
}
