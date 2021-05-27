package mycode

func maxProduct(words []string) int {
	masks := make([]int, len(words))

	for i, w := range words {
		for _, c := range w {
			masks[i] |= 1 << int(c-'a')
		}
	}

	ans := 0
	for i, m1 := range masks {
		for j := i + 1; j < len(masks); j++ {
			m2 := masks[j]
			if m1&m2 == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}

	return ans
}
