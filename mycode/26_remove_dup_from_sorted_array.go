package mycode

func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	} else {
		p := 1
		cur := 1
		for p < len(nums) {
			if nums[p] != nums[p-1] {
				nums[cur] = nums[p]
				cur++
			}
			p++
		}
		return cur
	}
}
