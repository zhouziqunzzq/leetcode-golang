package mycode

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func hasLeft(node *TreeNode) bool {
	return node.Left != nil
}

func hasRight(node *TreeNode) bool {
	return node.Right != nil
}

func dfs(node1 *TreeNode, node2 *TreeNode) bool {
	// recursion base
	if isLeaf(node1) && isLeaf(node2) {
		return true
	}
	if (isLeaf(node1)) && (!isLeaf(node2)) {
		return false
	}
	if (!isLeaf(node1)) && (isLeaf(node2)) {
		return false
	}
	// try to move down
	f1 := true
	if hasLeft(node1) {
		if hasLeft(node2) && node1.Left.Val == node2.Left.Val {
			// not flipped
			f1 = dfs(node1.Left, node2.Left)
		} else if hasRight(node2) && node1.Left.Val == node2.Right.Val {
			// flipped
			f1 = dfs(node1.Left, node2.Right)
		} else {
			f1 = false
		}
	}
	// early check
	if !f1 {
		return false
	}

	f2 := true
	if hasRight(node1) {
		if hasRight(node2) && node1.Right.Val == node2.Right.Val {
			// not flipped
			f2 = dfs(node1.Right, node2.Right)
		} else if hasLeft(node2) && node1.Right.Val == node2.Left.Val {
			// flipped
			f2 = dfs(node1.Right, node2.Left)
		} else {
			f1 = false
		}
	}
	return f1 && f2
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 != nil && root2 != nil {
		return dfs(root1, root2)
	} else if root1 == nil && root2 == nil {
		return true
	} else {
		return false
	}
}
