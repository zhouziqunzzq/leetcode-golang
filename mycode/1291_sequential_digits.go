package mycode

import "sort"

func sequentialDigits(low int, high int) []int {
	rst := make([]int, 0)

	for d := 1; d <= 9; d++ {
		sequentialDigitsHelper(d, low, high, &rst)
	}

	sort.Slice(rst, func(i, j int) bool {
		return rst[i] < rst[j]
	})

	return rst
}

func sequentialDigitsHelper(prev, low, high int, rst *[]int) {
	if prev >= low && prev <= high {
		*rst = append(*rst, prev)
	}

	lastDigit := prev % 10
	if lastDigit+1 >= 10 || prev > high {
		return
	}

	sequentialDigitsHelper(prev*10+lastDigit+1, low, high, rst)
}
