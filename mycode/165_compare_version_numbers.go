package mycode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	mul := 1
	if len(v1) < len(v2) {
		v1, v2 = v2, v1
		mul = -1
	}

	for i, _ := range v1 {
		n1, _ := strconv.Atoi(v1[i])
		n2 := 0
		if i < len(v2) {
			n2, _ = strconv.Atoi(v2[i])
		}
		if n1 > n2 {
			return 1 * mul
		} else if n1 < n2 {
			return -1 * mul
		}
	}
	return 0
}
