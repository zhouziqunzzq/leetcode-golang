package mycode

func numberOfSteps(num int) int {
	ans := 0
	for num != 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num--
		}
		ans++
	}
	return ans
}
