package mycode

func partitionLabels(S string) []int {
	lastIdx := make([]int, 26)
	ss := []rune(S)
	for i, c := range ss {
		lastIdx[c-'a'] = i
	}

	rst := make([]int, 0)
	curPartEnd, lastPartEnd := 0, -1
	for i, c := range ss {
		curPartEnd = maxInt(curPartEnd, lastIdx[c-'a'])
		if i == curPartEnd {
			rst = append(rst, curPartEnd-lastPartEnd)
			lastPartEnd = curPartEnd
		}
	}

	return rst
}
