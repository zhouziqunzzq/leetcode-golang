package mycode

import "strconv"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i, _ := range ans {
		if (i+1)%3 == 0 {
			ans[i] += "Fizz"
		}
		if (i+1)%5 == 0 {
			ans[i] += "Buzz"
		}
		if len(ans[i]) == 0 {
			ans[i] = strconv.FormatInt(int64(i+1), 10)
		}
	}
	return ans
}
