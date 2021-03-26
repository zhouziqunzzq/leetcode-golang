package mycode

// https://leetcode.com/problems/word-subsets/solution/
// Convert N_letter(a) >= N_letter(b_i) for each i to N_letter(a) >= max_i(N_letter(b_i))
// so we only need to save max N_letter for B
func wordSubsets(A []string, B []string) []string {
	maxCntB := make([]int, 26)
	for _, b := range B {
		cntB := wordSubsetsCount(b)
		for i := range cntB {
			maxCntB[i] = maxInt(maxCntB[i], cntB[i])
		}
	}

	ans := make([]string, 0)
	for _, a := range A {
		cntA := wordSubsetsCount(a)
		flag := true
		for i := range cntA {
			if cntA[i] < maxCntB[i] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, a)
		}
	}

	return ans
}

func wordSubsetsCount(s string) []int {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	return cnt
}
