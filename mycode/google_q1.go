package mycode

import "math"
import "strconv"

func solutionQ1(S string) string {
	// fmt.Fprintf(os.Stderr, "Tip: Use fmt.Fprintf(os.Stderr, \"\") to write debug messages on the output tab.")

	selection := []rune{'1', '2', '3'}
	return strconv.Itoa(findMinStrHelper([]rune(S), 0, selection))
}

func findMinStrHelper(s []rune, i int, selection []rune) int {
	// base case
	if i == len(s) {
		rst, _ := strconv.Atoi(string(s))
		return rst
	}

	if s[i] != '?' {
		// constraint checking
		if i != 0 && s[i-1] == s[i] {
			return math.MaxInt64
		} else {
			return findMinStrHelper(s, i+1, selection)
		}
	}

	// consider 1, 2, 3
	rst := math.MaxInt64
	for _, sl := range selection {
		// constraint checking
		if i != 0 && s[i-1] == sl {
			continue
		}
		s[i] = sl
		rst = minInt(rst, findMinStrHelper(s, i+1, selection))
		s[i] = '?'
	}
	return rst
}
