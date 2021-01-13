package mst

import "testing"

func TestKruskal(t *testing.T) {
	// adjecent matrix of  Graph
	G := [][]int{
		{0, 1, 3, 2},
		{1, 0, 0, 5},
		{3, 0, 0, 4},
		{2, 5, 4, 0},
	}

	t.Log(Kruskal(G))
}
