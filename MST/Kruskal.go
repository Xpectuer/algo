package mst

import "sort"

// Kruskal is a solution to MST problem
// G is an undirected Graph with weight
// A is edge set for spanning Tree
// pseudo code REF: Introduction to ALgorithm -
// O(ElogV) regardless of getEdges()
func Kruskal(G [][]int) [][]int {
	n := len(G)
	edges := getEdges(G)
	A := [][]int{}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[i][2]
	})

	// Disjoin Set Delaration
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	// end of initiation of Disjoin set

	var Find = func(x int) int {
		root := x
		for root != parent[root] {
			root = parent[root]
		}
		// path compression
		parent[x] = root
		return root

	}

	var Union = func(x, y int) {
		rootX, rootY := Find(x), Find(y)
		if rootX == rootY {
			return
		}
		parent[rootX] = rootY
		return
	}

	for _, e := range edges {
		if Find(e[0]) != Find(e[1]) {
			A = append(A, []int{e[0], e[1]})
			Union(e[0], e[1])
		}
	}
	return A
}

func getEdges(G [][]int) [][]int {
	n := len(G)
	edges := make([][]int, 0, n)
	for i := range G {
		for j := range G[0] {
			edges = append(edges, []int{i, j, G[i][j]})
		}
	}
	return edges
}
