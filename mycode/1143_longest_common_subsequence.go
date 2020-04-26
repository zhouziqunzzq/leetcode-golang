package mycode

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var opt [1000][1000]int

	// base case
	if text1[0] == text2[0] {
		opt[0][0] = 1
	}
	for i := 1; i < len(text1); i++ {
		if text1[i] == text2[0] || opt[i-1][0] == 1 {
			opt[i][0] = 1
		}
	}
	for j := 1; j < len(text2); j++ {
		if text1[0] == text2[j] || opt[0][j-1] == 1 {
			opt[0][j] = 1
		}
	}

	// fill the table
	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				opt[i][j] = opt[i-1][j-1] + 1
			} else {
				opt[i][j] = maxInt(opt[i][j-1], opt[i-1][j])
			}
		}
	}

	return opt[len(text1)-1][len(text2)-1]
}

func longestCommonSubsequenceFaster(text1 string, text2 string) int {
	opt := make([][]int, len(text1)+1)

	for i := range opt {
		opt[i] = make([]int, len(text2)+1)
	}

	for i1, c1 := range text1 {
		for i2, c2 := range text2 {
			if c1 == c2 {
				opt[i1+1][i2+1] = opt[i1][i2] + 1
			} else {
				max := opt[i1][i2+1]
				if opt[i1+1][i2] > max {
					max = opt[i1+1][i2]
				}
				opt[i1+1][i2+1] = max
			}
		}
	}

	return opt[len(text1)][len(text2)]
}
