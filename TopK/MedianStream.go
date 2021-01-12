package topk

import "container/heap"

type MedianFinder struct {
	// see IntHeap.go
	lo maxHeap
	hi minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	lo := make(maxHeap, 0)
	hi := make(minHeap, 0)

	heap.Init(&lo)
	heap.Init(&hi)
	return MedianFinder{lo, hi}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, num)
	heap.Push(&this.hi, heap.Pop(&this.lo))
	if this.lo.Len() < this.hi.Len() {
		heap.Push(&this.lo, heap.Pop(&this.hi))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		// odd quantity
		return float64(this.lo.Peek().(int))
	}
	return float64(this.lo.Peek().(int)+this.hi.Peek().(int)) * 0.5

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
