package mycode

// https://leetcode.com/problems/score-of-parentheses/solution/
// Solution 1: divide and conquer
// Solution 2: using stack
func scoreOfParentheses(S string) int {
	st := make([]int, 1)
	st[0] = 0

	for _, c := range S {
		if c == '(' {
			// go to a deeper level
			st = append(st, 0)
		} else {
			// return from a deeper level
			// add 2 * st[-1] to st[-2]
			lv2, lv1 := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-1]
			if lv2 == 0 {
				lv1 += 1
			} else {
				lv1 += 2 * lv2
			}
			st[len(st)-1] = lv1
		}
	}

	return st[0]
}

// Solution 3: counting "core"s
// core = ()
// the gist is that only cores provide actual value to the final answer,
// other parentheses are just multiplying core value by powers of 2.
func scoreOfParentheses3(S string) int {
	ans := 0
	balance := 0

	for i, c := range S {
		if c == '(' {
			balance++
		} else {
			balance--
			if S[i-1] == '(' {
				// core detected
				ans += 1 << balance
			}
		}
	}

	return ans
}
