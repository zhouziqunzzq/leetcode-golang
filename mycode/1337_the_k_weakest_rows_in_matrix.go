package mycode

import "sort"

type Row1337 struct {
	rid int
	cnt int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := make([]Row1337, len(mat))
	for i, r := range mat {
		cnt := 0
		for _, c := range r {
			if c == 0 {
				break
			}
			cnt++
		}
		rows[i] = Row1337{i, cnt}
	}

	sort.Slice(rows, func(i, j int) bool {
		if rows[i].cnt < rows[j].cnt {
			return true
		} else if rows[i].cnt == rows[j].cnt {
			return rows[i].rid < rows[j].rid
		} else {
			return false
		}
	})

	rst := make([]int, k)
	for i := range rst {
		rst[i] = rows[i].rid
	}
	return rst
}
