package mycode

// https://leetcode.com/problems/image-overlap/solution/
// using naive solution
func largestOverlap(A [][]int, B [][]int) int {
	rst := 0
	for yShift := 0; yShift < len(A); yShift++ {
		for xShift := 0; xShift < len(A); xShift++ {
			rst = maxInt(rst, largestOverlapHelper(xShift, yShift, A, B))
			rst = maxInt(rst, largestOverlapHelper(xShift, yShift, B, A))
		}
	}
	return rst
}

func largestOverlapHelper(xShift, yShift int, M, N [][]int) int {
	cnt := 0
	ny := 0
	for my := yShift; my < len(M); my++ {
		nx := 0
		for mx := xShift; mx < len(M); mx++ {
			if M[my][mx] == 1 && N[ny][nx] == 1 {
				cnt++
			}
			nx++
		}
		ny++
	}
	return cnt
}
