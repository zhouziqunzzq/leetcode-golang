package mycode

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	var mask, buf uint32
	subStrCnt := make(map[uint32]bool)
	desiredLen := 1 << k
	mask = (1 << k) - 1

	for i, c := range s {
		b := uint32(c - '0')
		buf = ((buf << 1) | b) & mask
		if i >= k-1 {
			subStrCnt[buf] = true
			if len(subStrCnt) == desiredLen {
				return true
			}
		}
	}

	return false
}
