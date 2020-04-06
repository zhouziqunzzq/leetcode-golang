package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else {
		l := len(nums)
		return &TreeNode{
			Val:   nums[l/2],
			Left:  createTree108(nums, 0, l/2),
			Right: createTree108(nums, l/2+1, l),
		}
	}
}

// [l, r)
func createTree108(nums []int, l int, r int) *TreeNode {
	if l == r {
		return nil
	} else {
		length := r - l
		m := l + length/2
		return &TreeNode{
			Val:   nums[m],
			Left:  createTree108(nums, l, m),
			Right: createTree108(nums, m+1, r),
		}
	}
}
