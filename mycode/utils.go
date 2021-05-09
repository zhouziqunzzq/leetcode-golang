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

func absInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Do a binary-search in a[beg:end] and return found index.
// Note that the returned index could be idx where a[idx] == target (when target is in a),
// or it could be idx where a[idx-1] < target and a[idx] > target (when target is not in a).
func intSlicesBinarySearch(a []int, beg, end, target int) int {
	l, r := beg, end
	m := 0
	for l < r {
		m = l + (r-l)/2
		if a[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func reverseInt(x int) int {
	ans := 0
	for x > 0 {
		ans = 10*ans + x%10
		x /= 10
	}
	return ans
}

func isPalindromeInt(x int) bool {
	return x == reverseInt(x)
}
