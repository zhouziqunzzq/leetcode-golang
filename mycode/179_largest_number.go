package mycode

import "strconv"
import "sort"
import "strings"

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	numStrs := make([]string, len(nums))
	for i, n := range nums {
		numStrs[i] = strconv.Itoa(n)
	}

	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	if numStrs[0] == "0" {
		return "0"
	} else {
		var sb strings.Builder
		for _, str := range numStrs {
			sb.WriteString(str)
		}
		return sb.String()
	}
}
