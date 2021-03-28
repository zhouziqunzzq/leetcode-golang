package mycode

import "strings"

// https://leetcode.com/problems/reconstruct-original-digits-from-english/discuss/91207/one-pass-O(n)-JAVA-Solution-Simple-and-Clear
// Count unique letter for each word of number
func originalDigits(s string) string {
	cnt := make([]int, 10)

	for _, c := range s {
		if c == 'z' {
			cnt[0]++
		}
		if c == 'w' {
			cnt[2]++
		}
		if c == 'u' {
			cnt[4]++
		}
		if c == 'x' {
			cnt[6]++
		}
		if c == 'g' {
			cnt[8]++
		}
		if c == 'o' {
			cnt[1]++ // 1, 0, 2, 4
		}
		if c == 'h' {
			cnt[3]++ // 3, 8
		}
		if c == 'f' {
			cnt[5]++ // 5, 4
		}
		if c == 's' {
			cnt[7]++ // 7, 6
		}
		if c == 'i' {
			cnt[9]++ // 9, 5, 6, 8
		}
	}

	cnt[1] -= cnt[0] + cnt[2] + cnt[4]
	cnt[3] -= cnt[8]
	cnt[5] -= cnt[4]
	cnt[7] -= cnt[6]
	cnt[9] -= cnt[5] + cnt[6] + cnt[8]

	sb := strings.Builder{}
	for i, c := range cnt {
		for j := 0; j < c; j++ {
			sb.WriteRune(rune('0' + i))
		}
	}

	return sb.String()
}
