package mycode

import "strings"

func findDuplicate(paths []string) [][]string {
	idx := make(map[string][]string) // file content -> list of files(paths)

	for _, path := range paths {
		t := strings.Split(path, " ")
		base := t[0] + "/"
		for _, file := range t[1:] {
			tt := strings.Split(file, "(")
			fullPath := base + tt[0]
			fileContent := tt[1][:len(tt[1])-1]
			if _, ok := idx[fileContent]; !ok {
				idx[fileContent] = make([]string, 0)
			}
			idx[fileContent] = append(idx[fileContent], fullPath)
		}
	}

	ans := make([][]string, 0)
	for _, l := range idx {
		if len(l) > 1 {
			ans = append(ans, l)
		}
	}

	return ans
}
