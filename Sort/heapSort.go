package mySort

func heapSort(A *[]int) {
	if len(*A) <= 1 {
		return
	}
	heapInit(A)

	// fmt.Println(*A)
	for k := len(*A) - 1; k > 0; k-- {
		(*A)[0], (*A)[k] = (*A)[k], (*A)[0]
		siftdown(A, k, 0)
	}
}

func heapInit(A *[]int) {
	n := len(*A)
	for i := n / 2; i >= 0; i-- {
		siftdown(A, n, i)
	}
}

// A: the heap
// n: the size of the heap
// i0: the index of the element to sift down
func siftdown(A *[]int, n, i0 int) {
	// for each node
	// layer0 / 1 node * h
	// layer1 / 2 node * (h-1)
	// layer2 / 2^2 node * (h-2)
	//...
	// Hence, T(n) = 1*h + 2*(h-1) + 2^2 * (h-2)
	// h = logN
	// O(h) = O(logn)
	i := i0
	for {
		j1 := i*2 + 1 // j1: left child
		// in case of overflow
		if j1 < 0 || j1 >= n {
			break
		}
		j := j1
		// j2: right child
		if j2 := j1 + 1; j2 < n && (*A)[j2] > (*A)[j1] {
			j = j2 // 2 * i + 2
		}
		// obey the definition of heap
		if (*A)[i] >= (*A)[j] {
			break
		}

		(*A)[i] ^= (*A)[j]
		(*A)[j] ^= (*A)[i]
		(*A)[i] ^= (*A)[j]
		// keep sifting down
		i = j
	}
}
