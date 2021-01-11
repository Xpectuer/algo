package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	A := []int{2, 3, 4, 1, 2, 5, 10, 2, 1}

	QuickSort(A[:], 0, len(A)-1)
	//sort.Ints(A)
	t.Log(A)

}
