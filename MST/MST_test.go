/*
 * @Author: your name
 * @Date: 2021-01-18 21:46:06
 * @LastEditTime: 2021-01-19 23:39:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /algo/MST/MST_test.go
 */
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

func TestPG(t *testing.T) {
	a := 1
	t.Log(a)

}
