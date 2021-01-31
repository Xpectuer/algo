/*
 * @Author: your name
 * @Date: 2021-01-18 21:34:28
 * @LastEditTime: 2021-01-31 11:04:46
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /algo/MST/Prim.go
 */
package mst

import (
	"container/heap"
	"fmt"
)

// Prim is a algorithm using greedy strategy to
// generate the MST
// returns edge set of MST
// G is a graph adjacent matrix | 0 represents not connected
// TODO:How can we trim the redundant output?

func Prim(G [][]int) [][]int {
	if len(G) == 0 || len(G[0]) == 0 {
		panic("The graph cannot be empty!")
	}

	// initializing
	var (
		visited int32
		queue   []int
		edges   PriorityQueue
	)

	queue = make([]int, 0, len(G))

	// high -> low: last -> first
	// select bit empty at first using bit packing
	// select the first

	// 11111
	visited = (1 << (len(G) + 1)) - 1
	// step1. select an arbitary node (the first one for example)
	heap.Init(&edges)
	queue = append(queue, 0)
	// set the no.0 as 0
	// 11110
	visited &= ^(1 << 0)

	// {{0,1},{1,2}}
	result := make([][]int, 0, len(G))
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for k := range G[i] {
			if G[i][k] != 0 {
				// push edges
				// may panic CBL in testing
				heap.Push(&edges, []int{i, k, G[i][k]})
				//fmt.Printf("%b\n", visited)
				// get the k bit, check visited
				if (visited>>k)&1 != 0 {
					queue = append(queue, k)
					visited &= ^(1 << k)
					fmt.Printf("%b\n", visited)
				}

			}
		}
		minEdge := heap.Pop(&edges).([]int)

		result = append(result, []int{minEdge[0], minEdge[1]})

	}
	return result

}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][2] < pq[j][2]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[0]
	*pq = old[1:]
	return x
}
