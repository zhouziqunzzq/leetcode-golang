package mycode

func findErrorNums(nums []int) []int {
	n := len(nums)
	cnt := make(map[int]int)

	dupElem := 0
	for _, v := range nums {
		if _, ok := cnt[v]; !ok {
			cnt[v] = 0
		}
		cnt[v]++
		if cnt[v] > 1 {
			dupElem = v
		}
	}

	missElem := 0
	for i := 1; i <= n; i++ {
		if _, ok := cnt[i]; !ok {
			missElem = i
			break
		}
	}

	return []int{dupElem, missElem}
}
