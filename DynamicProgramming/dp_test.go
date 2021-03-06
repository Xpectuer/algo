package dp

import "testing"

var G [][]int

func init() {
	G = [][]int{
		{0, 1, 4, 6},
		{0, 0, 0, 2},
		{0, 0, 0, 5},
		{0, 0, 0, 0},
	}
}

func TestBellmanFord(t *testing.T) {
	dist, err := BellmanFord(G, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(dist)
}
