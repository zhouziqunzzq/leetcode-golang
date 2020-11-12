package mycode

func solutionQ2(N int, A []int) string {
	//fmt.Fprintf(os.Stderr, "Tip: Use fmt.Fprintf(os.Stderr, \"\") to write debug messages on the output tab.")

	s := make([]rune, N)

	// fill s with '-'
	for i := range s {
		s[i] = '-'
	}

	// fill out '+'
	for _, pos := range A {
		s[pos] = '+'
	}

	return string(s)
}
