/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package mySort

import (
	"testing"
)

var (
	A []int
	n int
)

func init() {
	A = []int{0, 3, 4, 6, 1, 4, 8, 9, 2}
	n = len(A)
}
func TestHeap(t *testing.T) {

	heapSort(&A)
	t.Log(A)
}

func TestQuick(t *testing.T) {
	quickSort(0, n-1, A)
	t.Log(A)
}

func TestMerge(t *testing.T) {
	mergeSort(0, n-1, A)
	t.Log(A)
}

func TestInsert(t *testing.T) {
	//insertSort(A)
	bubbleSort(A)
	t.Log(A)
}
