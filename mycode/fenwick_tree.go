package mycode

// LSB(i) = ((i) & -(i))
// zeroes all the bits except the least significant one
type FenwickTree struct {
	len  int
	elem []int
}

func CreateFenwickTree(len int) *FenwickTree {
	return &FenwickTree{
		len:  len + 1,
		elem: make([]int, len+1),
	}
}

// increase point idx by 1 and update all its descendents down to the leaves
func (t *FenwickTree) Increase(idx int) {
	for idx < t.len {
		t.elem[idx]++
		idx += idx & -idx
	}
}

// get the prefix sum [1, idx] by traversing from node idx up to the root
func (t *FenwickTree) Get(idx int) int {
	rst := 0
	for idx > 0 {
		rst += t.elem[idx]
		idx -= idx & -idx
	}
	return rst
}
