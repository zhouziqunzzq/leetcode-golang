package mycode

// https://leetcode.com/problems/reach-a-number/solution/
// go over the math in solution...
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	k := 0
	for target > 0 {
		k++
		target -= k
	}

	if target%2 == 0 {
		return k
	} else {
		return k + 1 + k%2
	}
}
