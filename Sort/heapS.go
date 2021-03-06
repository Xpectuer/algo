package mySort

func heapSort1(A []int) {
	if len(A) == 0 {
		panic(" length error !")
	}
	n := len(A)
	heapify(A)

	for k := n - 1; k > 0; k-- {
		A[k], A[0] = A[0], A[k]
		siftdown1(A, k, 0)
	}
}

func heapify(A []int) {
	n := len(A)
	// initialize each non-leaves nodes
	for i := n / 2; i >= 0; i-- {
		siftdown1(A, n, i)
	}
}

func siftdown1(A []int, n, i0 int) {
	i := i0
	for {
		// left child
		j1 := i*2 + 1

		if j1 < 0 || j1 >= n {
			break
		}

		if j2 := j1 + 1; j2 < n && A[j2] > A[j1] {
			j1 = j2
		}
		if A[j1] <= A[i] {
			break
		}

		A[i], A[j1] = A[j1], A[i]

		i = j1
	}
}
