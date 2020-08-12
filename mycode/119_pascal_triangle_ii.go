package mycode

func getRow(rowIndex int) []int {
	rst := make([]int, rowIndex+1)
	rst[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			rst[j] += rst[j-1]
		}
	}
	return rst
}
