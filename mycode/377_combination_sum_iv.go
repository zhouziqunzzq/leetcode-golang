package mycode

// https://leetcode.com/problems/combination-sum-iv/discuss/85036/1ms-Java-DP-Solution-with-Detailed-Explanation
// dp
// dp[i] = sum(dp[i-nums[j]]) for j in 0...len(nums)-1 and i >= nums[j]
// dp[0] = 1
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i-n >= 0 {
				dp[i] += dp[i-n]
			}
		}
	}

	return dp[target]
}
