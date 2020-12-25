package mycode

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return make([]int, 0)
	}
	n := len(matrix[0])
	rst := make([]int, m*n)
	i := 0
	sweepUp := true

	for idxSum := 0; idxSum < m+n-1; idxSum++ {
		if sweepUp {
			for x := minInt(m-1, idxSum); x >= 0; x-- {
				y := idxSum - x
				if y >= n {
					break
				}
				rst[i] = matrix[x][y]
				i++
			}
		} else {
			for y := minInt(n-1, idxSum); y >= 0; y-- {
				x := idxSum - y
				if x >= m {
					break
				}
				rst[i] = matrix[x][y]
				i++
			}
		}

		sweepUp = !sweepUp
	}

	return rst
}
