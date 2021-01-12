package mycode

// https://leetcode.com/problems/create-sorted-array-through-instructions/discuss/927531/JavaC%2B%2BPython-Binary-Indexed-Tree
// Binary Indexed Tree (Fenwick Tree)
// https://en.wikipedia.org/wiki/Fenwick_tree
func createSortedArray(instructions []int) int {
	t := CreateFenwickTree(1e5)

	ans := 0
	for i, e := range instructions {
		ans = (ans + minInt(t.Get(e-1), i-t.Get(e))) % (1e9 + 7)
		t.Increase(e)
	}

	return ans
}
