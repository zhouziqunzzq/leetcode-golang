package mycode

func reorderedPowerOf2(N int) bool {
	cnt1 := reorderedPowerOf2CntHelper(N)

	for i := 0; i <= 31; i++ {
		cnt2 := reorderedPowerOf2CntHelper(1 << i)
		flag := true
		for i := 0; i < len(cnt1); i++ {
			if cnt1[i] != cnt2[i] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func reorderedPowerOf2CntHelper(n int) []int {
	cnt := make([]int, 10)
	for n != 0 {
		cnt[n%10]++
		n /= 10
	}
	return cnt
}
