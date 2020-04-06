package mycode

import "strconv"

func addBinary(a string, b string) string {
	carry := 0
	pa := len(a) - 1
	pb := len(b) - 1
	ans := ""

	for pa >= 0 || pb >= 0 {
		x, y := 0, 0
		if pa < 0 {
			x = 0
		} else {
			x = int(a[pa] - '0')
			pa--
		}
		if pb < 0 {
			y = 0
		} else {
			y = int(b[pb] - '0')
			pb--
		}

		ans = strconv.FormatInt(int64((x+y+carry)%2), 10) + ans
		if x+y+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}
