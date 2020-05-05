package mycode

func firstUniqChar(s string) int {
	var cnt1, cnt2, t int32 = 0, 0, 0

	for _, c := range s {
		t = 1 << (c - 'a')
		if t & cnt1 == 0 {
			cnt1 |= t
		} else if t & cnt2 == 0 {
			cnt2 |= t
		}
	}

	for i, c := range s {
		t = 1 << (c - 'a')
		if t & cnt2 == 0 {
			return i
		}
	}
	return -1
}
