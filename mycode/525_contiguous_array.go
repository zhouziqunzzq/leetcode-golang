package mycode

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 0
	cnt := 0
	cntToLastId := make(map[int]int)
	cntToLastId[0] = -1

	for i, v := range nums {
		if v == 0 {
			cnt--
		} else {
			cnt++
		}
		if lastId, ok := cntToLastId[cnt]; ok {
			if i - lastId > ans {
				ans = i - lastId
			}
		} else {
			cntToLastId[cnt] = i
		}
	}

	return ans
}
