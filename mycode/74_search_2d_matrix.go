package mycode

func searchMatrix74(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	// search for possible row id
	l, r := 0, m-1
	row := 0
	for l < r-1 {
		mid := l + (r-l)/2
		elem := matrix[mid][0]
		if target < elem {
			r = mid - 1
		} else if target > elem {
			l = mid
		} else {
			return true
		}
	}
	if l == r {
		row = l
	} else {
		if target >= matrix[r][0] {
			row = r
		} else {
			row = l
		}
	}

	// search in that row
	l, r = 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		elem := matrix[row][mid]
		if target < elem {
			r = mid - 1
		} else if target > elem {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
