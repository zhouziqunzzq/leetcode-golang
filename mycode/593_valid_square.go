package mycode

// https://leetcode.com/problems/valid-square/discuss/103442/C%2B%2B-3-lines-(unordered_set)
// Idea: distances consisting of any two vertices of a square can be divided into two
// types: edge(shorter) and diagonal(longer).
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d := make(map[int]bool)

	d[validSquareDistance(p1, p2)] = true
	d[validSquareDistance(p1, p3)] = true
	d[validSquareDistance(p1, p4)] = true
	d[validSquareDistance(p2, p3)] = true
	d[validSquareDistance(p2, p4)] = true
	d[validSquareDistance(p3, p4)] = true

	// corner case: all 0
	for k, _ := range d {
		if k == 0 {
			return false
		}
	}
	return len(d) == 2
}

func validSquareDistance(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
