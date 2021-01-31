package mySort

import (
	"testing"
)

func TestHeap(t *testing.T) {

	A := []int{0, 3, 4, 6, 1, 4, 8, 9, 2}
	heapSort(&A)
	t.Log(A)
}
