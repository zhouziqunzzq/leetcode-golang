package mycode

import "sort"

// https://leetcode.com/problems/sort-the-matrix-diagonally/discuss/489749/JavaPython-Straight-Forward
func diagonalSort(mat [][]int) [][]int {
	diagBuf := make(map[int][]int)

	for i, row := range mat {
		for j, v := range row {
			if _, ok := diagBuf[i-j]; !ok {
				diagBuf[i-j] = make([]int, 0)
			}
			diagBuf[i-j] = append(diagBuf[i-j], v)
		}
	}

	for _, d := range diagBuf {
		sort.Slice(d, func(i, j int) bool {
			return d[i] < d[j]
		})
	}

	for i, row := range mat {
		for j := range row {
			mat[i][j] = diagBuf[i-j][0]
			diagBuf[i-j] = diagBuf[i-j][1:]
		}
	}

	return mat
}
