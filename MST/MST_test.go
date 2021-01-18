package mst

import "testing"

var G [][]int

func init() {
	G = [][]int{
		{0, 1, 3, 2},
		{1, 0, 0, 5},
		{3, 0, 0, 4},
		{2, 5, 4, 0},
	}
}

func TestKruskal(t *testing.T) {
	// adjecent matrix of  Graph
	t.Log(Kruskal(G))
}
func TestPrim(t *testing.T) {
	// x := 1<<5 - 1
	// t.Logf("%b\n", x)
	// t.Logf("%b\n", x&^(1<<2))
	t.Log(Prim(G))

}
