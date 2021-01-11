package disjoinSet

// UnionFind something
type UnionFind struct {
	parent []int
	// weight to parent
	//weight []float64
}

// NewUnionFind something
func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{parent: make([]int, size)}
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		//uf.weight[i] = 1.0
	}
	return uf
}

// Union something
func (uf *UnionFind) Union(x int, y int, value float64) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.parent[rootX] = uf.parent[rootY]
	//uf.weight[rootX] = uf.weight[y] * value / uf.weight[x]
}

// Find something
func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		//origin := uf.parent[x]
		uf.parent[x] = uf.Find(uf.parent[x])
		//uf.weight[x] *= uf.weight[origin]
	}
	return uf.parent[x]
}

//
// func (uf *UnionFind) IsConnected(x int, y int) float64 {
// 	rootX, rootY := uf.Find(x), uf.Find(y)
// 	if rootY == rootX {
// 		return uf.weight[x] / uf.weight[y]
// 	}
// 	return -1.0
// }
