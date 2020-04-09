package mycode

// Can you solve it in O(N) time and O(1) space? Yes.
func backspaceCompare(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	backCnt := 0

	for {
		// backspace S if needed
		backCnt = 0
		for i >= 0 && S[i] == '#' {
			for i >= 0 && S[i] == '#' {
				i--
				backCnt++
			}
			for i >= 0 && backCnt > 0 {
				if S[i] != '#' {
					backCnt--
				} else {
					backCnt++
				}
				i--
			}
		}

		// backspace T if needed
		backCnt = 0
		for j >= 0 && T[j] == '#' {
			backCnt = 0
			for j >= 0 && T[j] == '#' {
				j--
				backCnt++
			}
			for j >= 0 && backCnt > 0 {
				if T[j] != '#' {
					backCnt--
				} else {
					backCnt++
				}
				j--
			}
		}

		if i == -1 && j == -1 {
			return true
		} else if i == -1 && j != -1 {
			return false
		} else if i != -1 && j == -1 {
			return false
		} else if S[i] != T[j] {
			return false
		} else {
			i--
			j--
		}
	}
}
