package mycode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	} else {
		nz := 0
		p := 0

		for p < len(nums) {
			if nums[p] != 0 {
				nums[nz], nums[p] = nums[p], nums[nz]
				nz++
			}
			p++
		}
	}
}
