package mycode

// https://leetcode.com/problems/partition-equal-subset-sum/solution/
// We use bottom-up dp here
// dp[n][subSetSum]
// dp[i][j] = true iff subset sum j can be obtained with nums[0...i-1]
// dp[i][j] = false o.w.
// so, dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
func canPartition(nums []int) bool {
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}
	if totalSum%2 != 0 {
		return false
	}

	subSetSum := totalSum / 2
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, subSetSum+1)
	}

	// base case
	dp[0][0] = true

	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= subSetSum; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][subSetSum]
}

// space compression version
func canPartition1D(nums []int) bool {
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}
	if totalSum%2 != 0 {
		return false
	}

	subSetSum := totalSum / 2
	dp := make([]bool, subSetSum+1)

	// base case
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := subSetSum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[subSetSum]
}
