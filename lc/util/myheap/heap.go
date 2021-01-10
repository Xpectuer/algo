package myheap

import "container/heap"

const d int = 2

type myInt []int

func (mi myInt) Len() int           { return len(mi) }
func (mi myInt) Swap(i, j int)      { mi[i], mi[j] = mi[j], mi[i] }
func (mi myInt) Less(i, j int) bool { return mi[i] < mi[j] }
func (mi *myInt) Push(x interface{}) {
	*mi = append(*mi, x.(int))
}
func (mi *myInt) Pop() interface{} {
	old := *mi
	n := len(old)
	x := old[n-1]
	*mi = old[0 : n-1]
	return x
}

func main() {
	h := &myInt{1, 2, 3, 4}

	heap.Push(h, 3)
}
