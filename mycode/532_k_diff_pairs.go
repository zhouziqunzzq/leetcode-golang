package mycode

func findPairs(nums []int, k int) int {
	cnt := make(map[int]int)
	for _, n := range nums {
		if _, ok := cnt[n]; ok {
			cnt[n]++
		} else {
			cnt[n] = 1
		}
	}

	ans := 0
	for n, cn := range cnt {
		if k > 0 {
			if _, ok := cnt[n+k]; ok {
				ans++
			}
		} else { // k == 0
			if cn >= 2 {
				ans++
			}
		}
	}

	return ans
}
