package mycode

// Simulation
func champagneTower(poured int, query_row int, query_glass int) float64 {
	t := make([][]float64, 101)
	t[0] = make([]float64, 1)
	t[0][0] = float64(poured)
	isOverflow := poured > 1

	for i := 0; i <= query_row && isOverflow; i++ {
		t[i+1] = make([]float64, i+2)
		isOverflow = false
		for j := 0; j < len(t[i]); j++ {
			if t[i][j] > 1.0 {
				isOverflow = true
				f := (t[i][j] - 1.0) / 2.0
				t[i][j] = 1.0
				t[i+1][j] += f
				t[i+1][j+1] += f
			}
		}
	}

	if t[query_row] == nil {
		return 0.0
	} else {
		return t[query_row][query_glass]
	}
}
