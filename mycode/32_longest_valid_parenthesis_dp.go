package mycode

func longestValidParenthesesDp(s string) int {
	if len(s) <= 1 {
		return 0
	}

	opt := make([]int, len(s))
	ans := 0

	// base cases
	opt[0] = 0
	if s[0] == '(' && s[1] == ')' {
		opt[1] = 2
		ans = 2
	}

	// fill the table
	for i := 2; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// ....()
				opt[i] = opt[i-2] + 2
			} else {
				// ....))
				if i - 1 - opt[i-1] >= 0 && s[i - 1 - opt[i-1]] == '(' {
					// ..((...))
					if i - 1 - opt[i-1] - 1 >= 0 {
						opt[i] = opt[i - 1 - opt[i-1] - 1] + opt[i-1] + 2
					} else {
						opt[i] = 0 + opt[i-1] + 2
					}

				}
			}

			if opt[i] > ans {
				ans = opt[i]
			}
		}
	}

	return ans
}
