package mycode

func findDuplicates(nums []int) []int {
	rst := make([]int, 0)
	for _, n := range nums {
		if n < 0 {
			n = -n
		}
		idx := n - 1
		if nums[idx] < 0 {
			rst = append(rst, n)
		} else {
			nums[idx] = -nums[idx]
		}
	}
	return rst
}
