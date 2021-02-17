package mySort

func insertSort(A []int) {
	for i := range A {
		tmp := A[i]
		j := i
		for j > 0 && A[j-1] > tmp {
			A[j] = A[j-1]
			j--
		}
		A[j] = tmp
	}
}

func bubbleSort(A []int) {
	n := len(A)
	for i := range A {
		for j := i; j < n; j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}
