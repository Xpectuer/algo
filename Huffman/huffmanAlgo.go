package huffman

import (
	"container/heap"
)

type huffmanNode struct {
	weight int

	char byte
	//0
	left *huffmanNode
	//1
	right *huffmanNode
}

// NewHuffmanNode somthing
func NewHuffmanNode(char byte, weight int, left *huffmanNode, right *huffmanNode) *huffmanNode {
	return &huffmanNode{weight, char, left, right}
}

func constructHuffman(forest []*huffmanNode) *huffmanNode {

	pq := make(PriorityQueue, 0)
	for _, e := range forest {
		pq.Push(e)
	}
	heap.Init(&pq)

	for pq.Len() != 1 {
		a := heap.Pop(&pq).(*huffmanNode)
		b := heap.Pop(&pq).(*huffmanNode)
		sum := NewHuffmanNode('$', a.weight+b.weight, a, b)
		heap.Push(&pq, sum)
	}
	return heap.Pop(&pq).(*huffmanNode)
}

// GetCode somthing call GetCode(root,0,&res)
func GetCode(root *huffmanNode, code int, res *map[byte]int) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		(*res)[root.char] = code
		return
	}
	// traverse the huffman tree
	code <<= 1
	GetCode(root.left, code, res)
	code = (code << 1) + 1
	GetCode(root.right, code, res)

}

// PriorityQueue something
type PriorityQueue []*huffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push something
func (pq *PriorityQueue) Push(x interface{}) {

	node := x.(*huffmanNode)
	*pq = append(*pq, node)
}

// Pop something
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return node
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *huffmanNode, weight int) {

	item.weight = weight
	heap.Fix(pq, item.weight)
}
