package mycode

func checkValidStringDp(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	var opt [105][105]bool

	// base case "*"
	for i := 0; i < l; i++ {
		opt[i][i] = s[i] == '*'
	}

	// base case "()", "(*", "*)", "**"
	for i := 0; i < l-1; i++ {
		j := i+1
		if (s[i] == '(' || s[i] == '*') && (s[j] == ')' || s[j] == '*') {
			opt[i][j] = true
		}
	}

	// fill the table
	for length := 3; length <= l; length++ {
		for i := 0; i <= l - length; i++ {
			j := i + length - 1

			// case 1
			if (s[i] == '(' || s[i] == '*') && (s[j] == ')' || s[j] == '*') && opt[i+1][j-1] {
				opt[i][j] = true
				continue
			}

			// case 2
			for k := i; k < j; k++ {
				if opt[i][k] && opt[k+1][j] {
					opt[i][j] = true
					break
				}
			}
		}
	}

	return opt[0][l-1]
}
