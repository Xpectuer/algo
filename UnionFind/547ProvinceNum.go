package disjoinSet

//func findCircleNum(isConnected [][]int) int {
//	n := len(isConnected)
//	count := n
//	uf := NewUnionFind(n)
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			x := isConnected[i][j]
//			if x == 1 {
//				uf.Union(i, j, &count)
//			}
//		}
//	}
//	return count
//}

//type UnionFind struct {
//	parent []int
//}
//
//func NewUnionFind(n int) *UnionFind {
//	uf := &UnionFind{parent: make([]int, n)}
//	for i := range uf.parent {
//		uf.parent[i] = i
//	}
//	return uf
//}
//
//func (uf *UnionFind) Find(x int) int {
//	if uf.parent[x] != x {
//		uf.parent[x] = uf.Find(uf.parent[x])
//	}
//	return uf.parent[x]
//}
//
//func (uf *UnionFind) Union(i int, j int, count *int) {
//	rootI := uf.Find(i)
//	rootJ := uf.Find(j)
//
//	if rootI != rootJ {
//		uf.parent[rootI] = uf.parent[rootJ]
//		*count--
//	}
//
//}
