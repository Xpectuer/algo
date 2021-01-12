package topk

import "container/heap"

// Origined from TopK problems, only 1 number (No.K least) required
// ref: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/go-ji-yu-zi-ji-shi-xian-de-zui-da-dui-de-ji-chu-sh/
func findLeastKHeap(A []int, k int) int {

	pq := make(maxHeap, 0)
	heap.Init(&pq)

	for i := 0; i < k; i++ {
		heap.Push(&pq, A[i])
	}

	for i := k; i < len(A); i++ {
		if pq.Len() >= k && A[i] < pq.Peek().(int) {
			heap.Pop(&pq)
			heap.Push(&pq, A[i])
		}

	}
	return pq.Peek().(int)
}
