/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package topk

import "testing"

func TestFindLeastKHeap(t *testing.T) {

	A := []int{2, 3, 4, 1, 6, 19, 2, 4, 30, 4}
	k := 6

	res := findLeastKHeap(A, k)
	t.Log(res)
	// required: 4
}

func TestFindTopKQuickSearch(t *testing.T) {
	A := []int{2, 3, 4, 1, 6, 19, 2, 4, 30, 4}
	//A := []int {1,2,3,2,2,2,2,2,2,2,2,2,1,1,1,1,1,1}
	k := 6

	// res := findTopKQuickSearch(A, k)
	res := bfprt(A, 0, len(A)-1, k)
	t.Log(res)
}
