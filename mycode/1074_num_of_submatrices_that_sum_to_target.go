package mycode

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	for i := range matrix {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			cnt = make(map[int]int)
			cnt[0] = 1
			cur := 0
			for k := 0; k < m; k++ {
				cur += matrix[k][j]
				if i > 0 {
					cur -= matrix[k][i-1]
				}
				if v, ok := cnt[cur-target]; ok {
					res += v
				}
				if v, ok := cnt[cur]; ok {
					cnt[cur] = v + 1
				} else {
					cnt[cur] = 1
				}
			}
		}
	}

	return res
}
