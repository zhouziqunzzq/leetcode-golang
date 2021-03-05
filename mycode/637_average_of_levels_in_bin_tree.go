package mycode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	sum := make([]int, 0)
	cnt := make([]int, 0)

	averageOfLevelsDFS(root, 0, &sum, &cnt)

	rst := make([]float64, len(sum))
	for i := range rst {
		rst[i] = float64(sum[i]) / float64(cnt[i])
	}

	return rst
}

func averageOfLevelsDFS(node *TreeNode, level int, sum, cnt *[]int) {
	if node == nil {
		return
	}

	if level >= len(*sum) {
		*sum = append(*sum, 0)
	}
	if level >= len(*cnt) {
		*cnt = append(*cnt, 0)
	}

	(*sum)[level] += node.Val
	(*cnt)[level]++

	for _, n := range []*TreeNode{node.Left, node.Right} {
		averageOfLevelsDFS(n, level+1, sum, cnt)
	}
}
