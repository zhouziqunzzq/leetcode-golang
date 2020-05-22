package mycode

import "sort"

func frequencySort(s string) string {
	cnt := make(map[rune]int)
	for _, r := range []rune(s) {
		if _, ok := cnt[r]; ok {
			cnt[r]++
		} else {
			cnt[r] = 1
		}
	}

	chars := make([]rune, len(cnt))
	charsIdx := 0
	for k, _ := range cnt {
		chars[charsIdx] = k
		charsIdx++
	}

	sort.Slice(chars, func(i, j int) bool {
		return cnt[chars[i]] > cnt[chars[j]]
	})

	ans := make([]rune, len(s))
	ansIdx := 0
	for _, c := range chars {
		for cnt[c] > 0 {
			ans[ansIdx] = c
			ansIdx++
			cnt[c]--
		}
	}

	return string(ans)
}
