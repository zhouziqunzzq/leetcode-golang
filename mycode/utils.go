package mycode

import "reflect"

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxInt3(x, y, z int) int {
	if x > y {
		if x > z {
			return x
		} else {
			return z
		}
	} else {
		if y > z {
			return y
		} else {
			return z
		}
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minInt3(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}

func reverseSlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
