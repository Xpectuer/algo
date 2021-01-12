package topk

import "fmt"

// Origined from TopK problems, only 1 number (No.K least) required
// ref https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/kuai-su-pai-xu-by-coder233/
func findTopKQuickSearch(A []int, k int) int {
	if k == 0 {
		panic("input k cannot be 0!")
	}
	start, end := 0, len(A)-1
	index := partition(A, start, end)
	for index != k-1 {
		// we find the topk least one
		// end when index = k-1
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(A, start, end)

	}
	return A[index]
}

// partition function,
// skip the normal(right&bigger than pivot)(left&smaller than pivot)
// if not normal , swap them (l, r)
// do not forget to restore the pivot
// (you should dig into the boundary situation)
func partition(A []int, l, r int) int {
	pivot := A[l]
	s, t := l, r+1
	for l < r {

		fmt.Printf("Before: %v\n", A[s:t])
		for l < r && A[r] >= pivot {
			r--
		}
		A[l] = A[r]
		for l < r && A[l] <= pivot {
			l++
		}
		A[r] = A[l]

		fmt.Printf("After: %v\n", A[s:t])
		fmt.Println("-------")
	}
	A[l] = pivot
	return l
}
