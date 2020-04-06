package mycode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	ans := 0
	findMaxDepth(root, 1, &ans)
	return ans
}

func findMaxDepth(v *TreeNode, nowDepth int, maxDepth *int) {
	if v == nil {
		nowDepth--
		if nowDepth > *maxDepth {
			*maxDepth = nowDepth
		}
		return
	} else {
		findMaxDepth(v.Left, nowDepth+1, maxDepth)
		findMaxDepth(v.Right, nowDepth+1, maxDepth)
	}
}
