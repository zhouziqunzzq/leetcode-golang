package mycode

func permuteUnique(nums []int) [][]int {
	rst := make([][]int, 0)
	buf := make([]int, 0)
	choices := make(map[int]int)

	for _, n := range nums {
		if _, ok := choices[n]; !ok {
			choices[n] = 1
		} else {
			choices[n]++
		}
	}

	permuteUniqueHelper(&buf, len(nums), choices, &rst)

	return rst
}

func permuteUniqueHelper(cur *[]int, n int, choices map[int]int, rst *[][]int) {
	// base case
	if len(*cur) == n {
		t := make([]int, n)
		for i := range t {
			t[i] = (*cur)[i]
		}
		*rst = append(*rst, t)
		return
	}

	// try different choices
	for i, cnt := range choices {
		if cnt == 0 {
			continue
		}

		*cur = append(*cur, i)
		choices[i]--

		permuteUniqueHelper(cur, n, choices, rst)

		*cur = (*cur)[:len(*cur)-1]
		choices[i]++
	}
}
