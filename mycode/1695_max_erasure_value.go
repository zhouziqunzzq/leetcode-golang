package mycode

// https://leetcode.com/problems/maximum-erasure-value/discuss/978552/Java-O(n)-Sliding-Window-%2B-HashSet
// Two ptrs (sliding window) + hashset
func maximumUniqueSubarray(nums []int) int {
	ans, sum := 0, 0
	set := make(map[int]bool)

	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		if _, ok := set[nums[j]]; !ok {
			sum += nums[j]
			ans = maxInt(ans, sum)
			set[nums[j]] = true
			j++
		} else {
			sum -= nums[i]
			delete(set, nums[i])
			i++
		}
	}

	return ans
}
