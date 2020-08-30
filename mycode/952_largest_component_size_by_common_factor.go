package mycode

// https://leetcode.com/problems/largest-component-size-by-common-factor/discuss/200959/Simplest-Solution-(Union-Find-only)-No-Prime-Calculation
func largestComponentSize(A []int) int {
	N := len(A)
	factorToNodeIdx := make(map[int]int)
	uf := CreateUnionFind(N)

	for i, a := range A {
		for j := 2; j*j <= a; j++ {
			if a%j == 0 {
				// j is a factor of a
				if idx, ok := factorToNodeIdx[j]; !ok {
					factorToNodeIdx[j] = i
				} else {
					uf.Union(i, idx)
				}
				// a/j is also a factor of a
				if idx, ok := factorToNodeIdx[a/j]; !ok {
					factorToNodeIdx[a/j] = i
				} else {
					uf.Union(i, idx)
				}
			}
		}
		// a is also a factor of a
		if idx, ok := factorToNodeIdx[a]; !ok {
			factorToNodeIdx[a] = i
		} else {
			uf.Union(i, idx)
		}
	}

	return uf.MaxSize
}
