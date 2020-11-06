package mycode

// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/solution/
// Solution: bin-search
func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, 1

	for _, n := range nums {
		if n > r {
			r = n
		}
	}

	for l <= r {
		m := l + (r-l)/2
		cur := smallestDivisorHelper(nums, m)
		if cur > threshold {
			// divisor is too small, go right
			l = m + 1
		} else {
			// divisor is too large, go left
			r = m - 1
		}
	}

	return l
}

func smallestDivisorHelper(nums []int, d int) int {
	sum := 0

	for _, n := range nums {
		sum += n / d
		if n%d != 0 {
			sum++
		}
	}

	return sum
}
