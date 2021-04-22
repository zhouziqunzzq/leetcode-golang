package mycode

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	dpPrev := make([]int, n)
	dpPrev[0] = triangle[0][0]
	dpCur := make([]int, n)

	for r := 1; r < n; r++ {
		for i := 0; i < r+1; i++ {
			a, b := math.MaxInt32, math.MaxInt32
			if i-1 >= 0 {
				a = dpPrev[i-1] + triangle[r][i]
			}
			if i < r {
				b = dpPrev[i] + triangle[r][i]
			}
			dpCur[i] = minInt(a, b)
		}
		dpPrev, dpCur = dpCur, make([]int, n)
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = minInt(ans, dpPrev[i])
	}

	return ans
}
