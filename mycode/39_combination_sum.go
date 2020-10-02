package mycode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	rst := make([][]int, 0)
	buf := make([]int, 0)
	combinationSumHelper(target, 0, candidates, &buf, &rst)
	return rst
}

func combinationSumHelper(remain, start int, candidates []int, buf *[]int, rst *[][]int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		tmp := make([]int, len(*buf))
		for i, n := range *buf {
			tmp[i] = n
		}
		*rst = append(*rst, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		*buf = append(*buf, candidates[i])
		combinationSumHelper(remain-candidates[i], i, candidates, buf, rst)
		*buf = (*buf)[:len(*buf)-1]
	}
}
