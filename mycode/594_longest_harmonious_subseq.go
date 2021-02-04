package mycode

func findLHS(nums []int) int {
	cnt := make(map[int]int)
	for _, n := range nums {
		if _, ok := cnt[n]; !ok {
			cnt[n] = 0
		}
		cnt[n]++
	}

	ans := 0
	for k, c := range cnt {
		// k-1 is automatically counted when current is k-1
		//if c1, ok := cnt[k-1]; ok {
		//	ans = maxInt(ans, c+c1)
		//}
		if c1, ok := cnt[k+1]; ok {
			ans = maxInt(ans, c+c1)
		}
	}

	return ans
}
