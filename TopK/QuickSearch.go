package topk

import "math/rand"

// Origined from TopK problems, only 1 number (No.K least) required
// ref https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/kuai-su-pai-xu-by-coder233/
func findTopKQuickSearch(A []int, k int) int {
	if k == 0 {
		panic("input k cannot be 0!")
	}
	start, end := 0, len(A)-1
	index := partition(A, start, end)
	// index := anotherPartition(A, start, end)
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
	// randomly select a index to improve the performance
	p := rand.Intn(r-l) + l
	A[l], A[p] = A[p], A[l]
	pivot := A[l]
	//s, t := l, r+1
	for l < r {

		//fmt.Printf("Before: %v\n", A[s:t])

		// > will occur dead looping
		for l < r && A[r] >= pivot {
			r--
		}
		A[l] = A[r]

		// < will occur dead looping
		for l < r && A[l] <= pivot {
			l++
		}
		A[r] = A[l]

		// fmt.Printf("After: %v\n", A[s:t])
		// fmt.Println("-------")
	}

	// redundant if do not use random selection
	A[l] = pivot
	return l
}

// anotherPartition is slower
// Reduce more entropy
// But return Sorted TopK result
func anotherPartition(A []int, l, r int) int {
	pivot := rand.Intn(r-l) + l

	counter := l
	A[r], A[pivot] = A[pivot], A[r]

	for i := l; i <= r; i++ {
		if A[i] < A[pivot] {
			A[i], A[counter] = A[counter], A[i]
			counter++
		}
	}
	A[counter], A[pivot] = A[pivot], A[counter]
	return counter
}

// 寻找第一五分位数（https://en.wikipedia.org/wiki/Quantile#:~:text=In%20statistics%20and%20probability%2C%20quantiles,the%20number%20of%20groups%20created.），作为parition的起始点
// copyleft: Blum、Floyd、Pratt、Rivest、Tarjan
func bfprt(A []int, l int, r int, k int) int {
	if k == 0 {
		panic("input k cannot be 0!")
	}
	start, end := 0, len(A)-1
	index := bfprtPartition(A, l, r)

	for index != k-1 {
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(A, start, end)
	}
	return A[index]

}

func bfprtPartition(A []int, l, r int) int {
	p := FindMid(A, l, r)

	A[l], A[p] = A[p], A[l]
	pivot := A[l]
	//s, t := l, r+1
	for l < r {
		//fmt.Printf("Before: %v\n", A[s:t])

		// > will occur dead looping
		for l < r && A[r] >= pivot {
			r--
		}
		A[l] = A[r]

		// < will occur dead looping
		for l < r && A[l] <= pivot {
			l++
		}
		A[r] = A[l]

		// fmt.Printf("After: %v\n", A[s:t])
		// fmt.Println("-------")
	}

	// not redundant here!!!
	A[l] = pivot
	return l
}

// FindMid is basically a Shell sort
// Sincerely Appriciate REF: https://zhuanlan.zhihu.com/p/31498036
func FindMid(A []int, l, r int) int {
	if l == r {
		return l
	}
	i, n := 0, 0
	for i := l; i < r-5; i += 5 {
		InsertSort(A, i, i+4)
		n = i - l
		A[l+n/5], A[i+2] = A[i+2], A[l+n/5]
	}
	// rest of the elements
	num := r - i + 1
	if num > 0 {
		InsertSort(A, i, i+num-1)
		n = i - l
		A[l+n/5], A[i+num/2] = A[i+num/2], A[l+n/5]

	}
	n /= 5
	if n == l {
		return l
	}
	return FindMid(A, l, l+n)
}

// InsertSort something
func InsertSort(A []int, l, r int) {

	for i := l+1; i <= r; i++ {
		if A[i-1] > A[i] {
			t := A[i]
			j := i
			for j > i && A[j-1] > t {
				A[j] = A[j-1]
				j--
			}
			A[j] = t
		}
	}
}
