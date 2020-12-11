package mycode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/discuss/27976/3-6-easy-lines-C%2B%2B-Java-Python-Ruby
// OP killed it! So elegant!
func removeDuplicates(nums []int) int {
	i := 0
	for _, n := range nums {
		if i < 2 || n != nums[i-2] {
			nums[i] = n
			i++
		}
	}
	return i
}
