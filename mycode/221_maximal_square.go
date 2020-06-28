package mycode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxL := 0
	opt := make([][]int, len(matrix))
	for i, _ := range opt {
		opt[i] = make([]int, len(matrix[0]))
	}

	// fill the table
	for i, m := range matrix {
		for j, v := range m {
			if v == '1' {
				if i == 0 || j == 0 {
					// base case
					opt[i][j] = 1
				} else {
					opt[i][j] = minInt3(opt[i-1][j], opt[i][j-1], opt[i-1][j-1]) + 1
				}

				// update global maximum
				if opt[i][j] > maxL {
					maxL = opt[i][j]
				}
			}
		}
	}

	return maxL * maxL
}

func maximalSquareBetter(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxL := 0
	opt := make([]int, len(matrix[0]))

	// fill the table
	prev := 0
	for i, m := range matrix {
		for j, v := range m {
			tmp := opt[j]
			if v == '1' {
				if i == 0 || j == 0 {
					// base case
					opt[j] = 1
				} else {
					opt[j] = minInt3(opt[j-1], opt[j], prev) + 1
				}

				// update global maximum
				if opt[j] > maxL {
					maxL = opt[j]
				}
			} else {
				opt[j] = 0
			}
			prev = tmp
		}
	}

	return maxL * maxL
}
