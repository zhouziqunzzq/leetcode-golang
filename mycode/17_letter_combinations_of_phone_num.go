package mycode

func letterCombinations(digits string) []string {
	buf := make([]rune, 5)
	rst := make([]string, 0)
	vocab := make(map[int]string)
	vocab[2] = "abc"
	vocab[3] = "def"
	vocab[4] = "ghi"
	vocab[5] = "jkl"
	vocab[6] = "mno"
	vocab[7] = "pqrs"
	vocab[8] = "tuv"
	vocab[9] = "wxyz"

	letterCombinationsDFS(digits, 0, buf, &rst, vocab)

	return rst
}

func letterCombinationsDFS(digits string, i int, buf []rune, rst *[]string, vocab map[int]string) {
	if i >= len(digits) {
		if i > 0 {
			t := make([]rune, i)
			for j := 0; j < i; j++ {
				t[j] = buf[j]
			}
			*rst = append(*rst, string(t))
		}
		return
	}

	d := int(digits[i] - '0')
	for _, c := range vocab[d] {
		buf[i] = c
		letterCombinationsDFS(digits, i+1, buf, rst, vocab)
	}
}
