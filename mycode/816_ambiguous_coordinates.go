package mycode

import "fmt"

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	ans := make([]string, 0)
	for i := 1; i < len(s); i++ {
		l := ambiguousCoordinatesHelper(s[:i])
		r := ambiguousCoordinatesHelper(s[i:])
		for _, lSeg := range l {
			for _, rSeg := range r {
				ans = append(ans, fmt.Sprintf("(%s, %s)", lSeg, rSeg))
			}
		}
	}
	return ans
}

func ambiguousCoordinatesHelper(seg string) []string {
	rst := make([]string, 0)
	n := len(seg)
	for d := 1; d <= n; d++ {
		left := seg[:d]
		right := seg[d:]
		if (left == "0" || left[0] != '0') && (len(right) == 0 || right[len(right)-1] != '0') {
			t := left
			if d != n {
				t += "." + right
			}
			rst = append(rst, t)
		}
	}
	return rst
}
