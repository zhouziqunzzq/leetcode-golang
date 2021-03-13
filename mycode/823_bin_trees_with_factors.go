package mycode

import "sort"

// https://leetcode.com/problems/binary-trees-with-factors/solution/
// DP: dp[i] = the number of ways to have a root node with value arr[i]
func numFactoredBinaryTrees(arr []int) int {
	const mod = 1_000_000_007
	n := len(arr)
	sort.Ints(arr)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	valueToIdx := make(map[int]int)
	for i, v := range arr {
		valueToIdx[v] = i
	}

	for i := 0; i < n; i++ { // arr[i] as root
		for j := 0; j < i; j++ { // arr[j] as left child
			if arr[i]%arr[j] == 0 {
				right := arr[i] / arr[j]
				if idx, ok := valueToIdx[right]; ok {
					dp[i] = (dp[i] + dp[j]*dp[idx]) % mod
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}

	return ans
}
