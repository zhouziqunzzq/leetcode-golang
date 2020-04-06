package mycode

import "strconv"

func groupAnagrams(strs []string) [][]string {
	hashToId := make(map[string]int)
	myHash := make([]int, 26)
	for i := 0; i < 26; i++ {
		myHash[i] = 0
	}

	rst := make([][]string, 0)

	for _, s := range strs {
		for i := 0; i < 26; i++ {
			myHash[i] = 0
		}

		for _, c := range s {
			myHash[c - 'a']++
		}

		hashStr := ""
		for _, h := range myHash {
			hashStr += strconv.Itoa(h)
		}

		if v, ok := hashToId[hashStr]; ok {
			rst[v] = append(rst[v], s)
		} else {
			newId := len(rst)
			hashToId[hashStr] = newId
			rst = append(rst, make([]string, 0))
			rst[newId] = append(rst[newId], s)
		}
	}

	return rst
}
