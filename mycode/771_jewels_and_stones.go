package mycode

// space usage can be optimized using bitmap instead of map...
func numJewelsInStones(J string, S string) int {
	ans := 0

	jewelManual := make(map[rune]bool)
	for _, j := range J {
		jewelManual[j] = true
	}

	for _, s := range S {
		if _, ok := jewelManual[s]; ok {
			ans++
		}
	}

	return ans
}
