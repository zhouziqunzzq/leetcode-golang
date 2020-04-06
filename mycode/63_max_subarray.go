package mycode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else {
		sum := nums[0]
		maxSum := nums[0]
		for i := 1; i < len(nums); i++ {
			if sum <= 0 {
				sum = nums[i]
			} else {
				sum += nums[i]
			}
			//maxSum = int(math.Max(float64(maxSum), float64(sum)))
			if sum > maxSum {
				maxSum = sum
			}
		}
		return maxSum
	}
}
