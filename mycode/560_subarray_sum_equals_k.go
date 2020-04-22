package mycode

// Solution 1: O(n^2) scanning.
func subarraySum(nums []int, k int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}

	return ans
}

// Solution 2: O(n) using hash table.
func subarraySumHash(nums []int, k int) int {
	ans := 0
	sum := 0
	sumCnt := make(map[int]int)
	sumCnt[0] = 1

	for _, v := range nums {
		sum += v

		if v, ok := sumCnt[sum-k]; ok {
			ans += v
		}

		if _, ok := sumCnt[sum]; ok {
			sumCnt[sum]++
		} else {
			sumCnt[sum] = 1
		}
	}

	return ans
}
