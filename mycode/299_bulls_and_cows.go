package mycode

import "fmt"

func getHint(secret string, guess string) string {
	a, b := 0, 0
	cnt := make([]int, 10)
	s := []rune(secret)
	g := []rune(guess)

	for _, c := range s {
		cnt[c-'0']++
	}

	for i, c := range g {
		if c == s[i] {
			a++
			cnt[c-'0']--
		}
	}

	for i, c := range g {
		if c != s[i] && cnt[c-'0'] > 0 {
			b++
			cnt[c-'0']--
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}
