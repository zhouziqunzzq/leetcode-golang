package mycode

func countBits(num int) []int {
	ans := make([]int, num+1)
	cnt := 1

	for j := 1; j <= num; j <<= 1 {
		for i := 0; i < j && cnt <= num; i++ {
			ans[j+i] = ans[i] + 1
			cnt++
		}
	}

	return ans
}
