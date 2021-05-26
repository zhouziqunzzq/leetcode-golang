package mycode

func minPartitions(n string) int {
	ans := 0
	for _, c := range n {
		ans = maxInt(ans, int(c-'0'))
	}
	return ans
}
