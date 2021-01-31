package mySort

import "fmt"

func heapSort(A *[]int) {
	if len(*A) <= 1 {
		return
	}
	heapInit(A)
	k := len(*A) - 1

	fmt.Println(*A)
	for k > 0 {
		(*A)[0], (*A)[k] = (*A)[k], (*A)[0]
		k--
		siftdown(A, k+1, 0)
	}
}

func heapInit(A *[]int) {
	n := len(*A)
	for i := n / 2; i >= 0; i-- {
		siftdown(A, n, i)
	}
}

func siftdown(A *[]int, n, i0 int) {
	// for each node
	// layer0 / 1 node * h
	// layer1 / 2 node * (h-1)
	// layer2 / 2^2 node * (h-2)
	//...
	// Hence, T(n) = 1*h + 2*(h-1) + 2^2 * (h-2)
	// h = logN
	// O(n)
	i := i0
	for {
		j1 := i*2 + 1 // left child
		if j1 < 0 || j1 >= n {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && (*A)[j2] > (*A)[j1] {
			j = j2 // 2 * i + 2
		}
		if (*A)[i] >= (*A)[j] {
			break
		}

		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i = j
	}
}
