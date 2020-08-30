package mycode

type UnionFind struct {
	Parent  []int
	Size    []int
	MaxSize int
}

func CreateUnionFind(N int) UnionFind {
	uf := UnionFind{
		Parent:  make([]int, N),
		Size:    make([]int, N),
		MaxSize: 1,
	}
	for i := 0; i < N; i++ {
		uf.Parent[i] = i
		uf.Size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if x == uf.Parent[x] {
		return x
	} else {
		return uf.Find(uf.Parent[x])
	}
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		uf.Parent[rootX] = rootY
		uf.Size[rootY] += uf.Size[rootX]
		uf.MaxSize = maxInt(uf.MaxSize, uf.Size[rootY])
	}
}
