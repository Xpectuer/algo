package mySort

import "fmt"

func quickSort(l, r int, A []int) {
	if l >= r {
		return
	}

	p := partition(l, r, A)
	fmt.Println(p)
	quickSort(l, p-1, A)
	quickSort(p+1, r, A)

}

func partition(l, r int, A []int) int {
	pivot := r
	counter := l
	for i := l; i < r; i++ {
		if (A)[i] < (A)[pivot] {
			(A)[i], (A)[counter] = (A)[counter], (A)[i]
			counter++
		}
	}
	(A)[pivot], (A)[counter] = (A)[counter], (A)[pivot]
	return counter
}
