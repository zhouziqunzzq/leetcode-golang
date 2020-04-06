package mycode

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		s := []rune(countAndSay(n-1))
		num := s[0]
		cnt := 0
		var rst string

		for _, c := range s {
			if num == c {
				cnt++
			} else {
				rst += strconv.Itoa(cnt) + string(num)
				num = c
				cnt = 1
			}
		}

		rst += strconv.Itoa(cnt) + string(num)

		return rst
	}
}
