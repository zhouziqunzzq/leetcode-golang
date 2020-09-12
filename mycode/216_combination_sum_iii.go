package mycode

func combinationSum3(k int, n int) [][]int {
	rst := make([][]int, 0)
	buf := make([]int, k)
	var flags [10]bool // flags[i] = true <=> digit i was used
	combinationSum3Helper(0, 0, k, n, 1, flags, buf, &rst)
	return rst
}

func combinationSum3Helper(curK, curN, k, n, nextI int, flags [10]bool, buf []int, rst *[][]int) {
	if curK == k {
		if curN == n {
			s := make([]int, len(buf))
			for i, b := range buf {
				s[i] = b
			}
			*rst = append(*rst, s)
		}
		return
	}

	for i := nextI; i <= 9; i++ {
		if flags[i] || curN+i > n {
			continue
		}
		buf[curK] = i
		flags[i] = true
		combinationSum3Helper(curK+1, curN+i, k, n, i+1, flags, buf, rst)
		flags[i] = false
	}
}
