package mycode

func longestPalindrome409(s string) int {
	var cnt [52]int

	for _, c := range []rune(s) {
		if c >= 'a' && c <= 'z' {
			cnt[c-'a']++
		} else {
			cnt[c-'A'+26]++
		}
	}

	hasOdd := 0
	rst := 0
	for _, n := range cnt {
		if n%2 == 1 {
			hasOdd = 1
			n -= 1
		}
		rst += n
	}
	rst += hasOdd

	return rst
}
