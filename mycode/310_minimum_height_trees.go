package mycode

// https://leetcode.com/problems/minimum-height-trees/solution/
// Topological sort
func findMinHeightTrees(n int, edges [][]int) []int {
	leaves := make([]int, 0)
	if n <= 2 {
		for i := 0; i < n; i++ {
			leaves = append(leaves, i)
		}
		return leaves
	}

	// build adj list to store graph
	neighbors := make([]map[int]bool, n)
	for i := range neighbors {
		neighbors[i] = make(map[int]bool)
	}
	for _, e := range edges {
		neighbors[e[0]][e[1]] = true
		neighbors[e[1]][e[0]] = true
	}

	// init leaves
	for i, n := range neighbors {
		if len(n) == 1 {
			leaves = append(leaves, i)
		}
	}

	remainingNodes := n
	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		newLeaves := make([]int, 0)

		for _, leaf := range leaves {
			for neighbor := range neighbors[leaf] {
				delete(neighbors[neighbor], leaf)
				if len(neighbors[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}
