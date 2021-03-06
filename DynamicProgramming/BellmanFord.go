package dp

import (
	"errors"
	"fmt"
	"math"
)

// A Bellman-Ford solved the shortest path problem of
//  single sourced graph

// BellmanFord is the dp algorithm,
// input a adjacent Matrix and source point.
func BellmanFord(G [][]int, src int) ([]int, error) {

	if src < 0 || src >= len(G) {
		return nil, errors.New("Invalid Source")
	}
	dist := make([]int, len(G))
	// Step 1: Initialize distances from src to all
	// other vertices
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	E := getEdges(G)
	// Step 2: Relax all edges |V| -1 times.
	// A simple path (that is path does not have any repeated vertices)
	// shortest path from src to any other vertex can have at-most |V| - 1
	// edges
	for i := 1; i <= len(G)-1; i++ {
		for j := 0; j < len(E); j++ {
			u := E[j][0]
			v := E[j][1]
			weight := E[j][2]
			if dist[u] != math.MaxInt32 && dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
			}
		}
	}

	// Step3 : Check for negative-weight cycles. The above step
	// guarantees shortest distance if graph doesn't contain
	// negative weight cycle. If we get a shorter path, then there is a cycle
	for i := 0; i < len(E); i++ {
		u := E[i][0]
		v := E[i][1]
		weight := E[i][2]
		if dist[u] != math.MaxInt32 && dist[u]+weight < dist[v] {
			fmt.Println(E[i])
			return nil, errors.New("The Graph Contains a negative cycle")
		}
	}

	return dist, nil
}

// return: {src, dest, distance}
func getEdges(G [][]int) [][]int {
	E := make([][]int, 0, len(G))
	for i := range G {
		for j := range G[0] {
			triplet := []int{i, j, G[i][j]}
			E = append(E, triplet)
		}
	}
	return E
}
