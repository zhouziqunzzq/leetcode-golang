package mycode

func findTheDifference(s string, t string) byte {
	var cnt [26]int
	n := len(s)

	for i := 0; i < n; i++ {
		cnt[s[i]-'a']--
		cnt[t[i]-'a']++
	}
	cnt[t[n]-'a']++

	for i, c := range cnt {
		if c > 0 {
			return byte('a' + i)
		}
	}
	return 'a'
}
