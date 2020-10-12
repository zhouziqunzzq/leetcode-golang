package mycode

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	diffIdx := make([]int, 0, 3)
	cnt := make([]int, 26)
	hasPair := false

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diffIdx = append(diffIdx, i)
		} else {
			cnt[A[i]-'a']++
			if cnt[A[i]-'a'] >= 2 {
				hasPair = true
			}
		}
		if len(diffIdx) > 2 {
			return false
		}

	}

	if len(diffIdx) == 0 {
		return hasPair
	}

	if len(diffIdx) == 2 {
		i, j := diffIdx[0], diffIdx[1]
		return A[i] == B[j] && A[j] == B[i]
	}

	return false
}
