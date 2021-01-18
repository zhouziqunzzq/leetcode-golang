package mycode

func maxOperations(nums []int, k int) int {
	cnt := make(map[int]int)

	for _, n := range nums {
		if _, ok := cnt[n]; !ok {
			cnt[n] = 0
		}
		cnt[n]++
	}

	ans := 0
	for _, n := range nums {
		m := k - n
		if n != m {
			p := minInt(cnt[n], cnt[m])
			ans += p
			cnt[n] -= p
			cnt[m] -= p
		} else {
			p := cnt[n] / 2
			ans += p
			cnt[n] -= 2 * p
		}
	}

	return ans
}
