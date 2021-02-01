package mycode

func hammingWeight(num uint32) int {
	ans := 0

	for num != 0 {
		ans++
		num &= num - 1
	}

	return ans
}
